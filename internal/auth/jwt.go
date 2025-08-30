package auth

import (
	"fmt"
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte
var ServerVersion int64

// Init генерирует новый секрет и обновляет версию сервера
func Init() {
	JWTSecret = fmt.Appendf(nil, "secret_%d", time.Now().UnixNano())
	ServerVersion = time.Now().UnixNano()
}

// CreateToken создает JWT с версией сервера
func CreateToken(user_id, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24 * 30).Unix(),
		"sv":      ServerVersion, // версия сервера
	})

	tokenString, err := token.SignedString(JWTSecret)
	if err != nil {
		return "", errors.New("JWT token can't be created")
	}

	return tokenString, nil
}

// VerifyToken проверяет JWT и версию сервера
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid JWT token")
	}

	claims := token.Claims.(jwt.MapClaims)

	// Проверяем версию сервера
	svFloat, ok := claims["sv"].(float64)
	if !ok || int64(svFloat) != ServerVersion {
		return nil, errors.New("Token expired due to server restart")
	}

	return claims, nil
}
