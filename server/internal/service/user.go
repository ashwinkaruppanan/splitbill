package service

import (
	"context"
	"database/sql"
	"net/http"
	"splitbill/internal/model"
	"splitbill/utils"
	"time"
)

type userService struct {
	repository  model.UserRepositoryInterface
	mailService *utils.MailService
}

func NewUserService(repository model.UserRepositoryInterface, mailService *utils.MailService) model.UserServiceInterface {
	return &userService{
		repository,
		mailService,
	}
}

func (s *userService) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, &utils.AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
	}

	user.Password = hashedPassword
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	// err = s.mailService.MailTo(user.Email, "SplitBill", "Account created successfully.")
	// if err != nil {
	// 	return &utils.AppError{Code: http.StatusInternalServerError, Message: "internal server errror"}
	// }

	res, err := s.repository.CreateUser(ctx, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &utils.AppError{Code: http.StatusConflict, Message: "email already exist"}
		}
		return nil, &utils.AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
	}

	res.Password = ""
	return res, nil
}

func (s *userService) LoginUser(ctx context.Context, user *model.User) (*model.User, error) {

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	res, err := s.repository.GetUserByEmail(ctx, user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, &utils.AppError{Code: http.StatusBadRequest, Message: "wrong email"}
		}
		return nil, &utils.AppError{Code: http.StatusInternalServerError, Message: "internal server error"}
	}

	err = utils.VerifyPassword(user.Password, res.Password)
	if err != nil {
		return nil, &utils.AppError{Code: http.StatusBadRequest, Message: "wrong password"}
	}

	res.Password = ""
	return res, nil
}
