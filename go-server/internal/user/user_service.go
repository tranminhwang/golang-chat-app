package user

import (
	"context"
	"fmt"
	"go-server/utils"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}
	r, err := s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserResponse{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

func (s *service) UserLogin(c context.Context, req *LoginRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	user, _ := s.Repository.UserLogin(ctx, req.Email)
	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	err := utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	return &CreateUserResponse{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Email:    user.Email,
	}, nil
}
