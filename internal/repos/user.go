package repos

import (
	"context"
	"errors"
	"fmt"
	"auth-service/internal/auth"
	"auth-service/internal/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUserInDB(pool *pgxpool.Pool, user *models.User) error {
	ctx := context.Background()

	transaction, err := pool.Begin(ctx)
    if err != nil {
        return err
    }

    // Откатим транзакцию, если произойдет паника или ошибка
    defer func() {
        if p := recover(); p != nil {
            transaction.Rollback(ctx)
            panic(p) // повторно вызываем панику после отката
        } else if err != nil {
            transaction.Rollback(ctx)
        } else {
            err = transaction.Commit(ctx) // коммитим, если ошибок нет
        }
    }()

	// Проверяем, есть ли уже такой пользователь
	var exists bool
	err = transaction.QueryRow(ctx,"SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", user.Username).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("User already exists")
	}

	// Хэшируем пароль
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	user.PasswordHash = hashedPassword

	// Вставляем нового пользователя
	err = pool.QueryRow(context.Background(),
		"INSERT INTO users (username, password_hash, role) VALUES ($1, $2, $3) RETURNING id, created_at",
		user.Username, hashedPassword, user.Role).Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}

	// Чистим пароль для безопасности
	user.Password = ""

	return nil
}

func LoginUser(pool *pgxpool.Pool, user *models.User) (string, error) {
	// Check if the user exists
	var storedPasswordHash, role string
	err := pool.QueryRow(context.Background(),
		"SELECT id, password_hash, role FROM users WHERE username = $1",
		user.Username).Scan(&user.ID, &storedPasswordHash, &role)

	// Объединяем ошибки: либо пользователь не найден, либо пароль неверный
	if err != nil || !auth.CheckPasswordHash(user.Password, storedPasswordHash) {
		return "", errors.New("Invalid username or password")
	}

	user.Role = role

	// Создаём токен
	token, err := auth.CreateToken(fmt.Sprintf("%d", user.ID), user.Role)
	if err != nil {
		return "", errors.New("Failed to create JWT")
	}

	return token, nil
}
