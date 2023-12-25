package model

import (
	"context"
	"time"
)

type User struct {
	User_id    int32     `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type UserRepositoryInterface interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type UserServiceInterface interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	LoginUser(ctx context.Context, user *User) (*User, error)
}
