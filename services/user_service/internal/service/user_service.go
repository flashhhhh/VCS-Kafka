package service

import (
	"context"
	"user_service/internal/domain"
	"user_service/internal/repository"
)

type UserService interface {
	CreateUser(ctx context.Context, username, password, name string) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, username, password, name string) error {
	user := &domain.User{
		Username: username,
		Password: password,
		Name:     name,
	}

	err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}