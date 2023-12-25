package repository

import (
	"context"
	"database/sql"
	"splitbill/internal/model"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) model.UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	var inserted_id int

	query := "INSERT INTO users (name, email, password, created_at, updated_at) SELECT $1, $2, $3, $4, $5 WHERE NOT EXISTS (SELECT 1 FROM users WHERE email = $6) RETURNING user_id;"

	err := u.db.QueryRowContext(ctx, query, user.Name, user.Email, user.Password, user.Created_at, user.Updated_at, user.Email).Scan(&inserted_id)
	if err != nil {
		return nil, err
	}
	user.User_id = int32(inserted_id)
	return user, nil
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	query := "SELECT * FROM users WHERE email = $1"

	err := u.db.QueryRowContext(ctx, query, email).Scan(&user.User_id, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
