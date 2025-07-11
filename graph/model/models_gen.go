// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type AddTraining struct {
	Name     string `json:"name"`
	Duration string `json:"duration"`
	Date     string `json:"date"`
	Notes    string `json:"notes"`
	Type     int32  `json:"type"`
}

type AuthPayload struct {
	Token string `json:"token"`
	User  *User  `json:"user"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Mutation struct {
}

type Query struct {
}

type RegisterInput struct {
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type SearchTrainings struct {
	Name   *string `json:"name,omitempty"`
	Type   *int32  `json:"type,omitempty"`
	Limit  int32   `json:"limit"`
	Offset int32   `json:"offset"`
}

type Training struct {
	ID        string     `json:"id"`
	UserID    string     `json:"user_id"`
	Name      string     `json:"name"`
	Duration  string     `json:"duration"`
	Date      string     `json:"date"`
	Notes     string     `json:"notes"`
	Type      int32      `json:"type"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

type User struct {
	ID         string     `json:"id"`
	Email      string     `json:"email"`
	Login      string     `json:"login"`
	Password   string     `json:"password"`
	Role       string     `json:"role"`
	TelegramID *string    `json:"telegram_id,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
}
