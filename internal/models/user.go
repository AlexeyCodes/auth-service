package models

import (
	"time"
)

// user struct
type User struct {
	ID           int       `db:"id" json:"id"`
	Username     string    `db:"username" json:"username" binding:"required,min=1"`
	PasswordHash string    `db:"password_hash" json:"-"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	Role         string    `db:"role" json:"role"`

	Password string `db:"-" json:"password" binding:"required,min=6"`
}
