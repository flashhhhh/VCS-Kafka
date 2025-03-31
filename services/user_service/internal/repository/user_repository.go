package repository

import (
	"context"
	"user_service/internal/domain"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
}

type userRepository struct {
	db *gorm.DB
	redis *redis.Client
}

func NewUserRepository(db *gorm.DB, redis *redis.Client) UserRepository {
	return &userRepository{
		db:    db,
		redis: redis,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *domain.User) error {
	err := r.db.Create(user).Error
	return err
}