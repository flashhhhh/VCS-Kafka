package repository

import (
	"order_service/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) CreateOrder(orders []domain.Order) error {
	return r.db.Create(&orders).Error
}

func (r *OrderRepository) GetMaxID() (int, error) {
	var maxID int
	err := r.db.Model(&domain.Order{}).Select("MAX(id)").Scan(&maxID).Error
	if err != nil {
		return 0, err
	}
	return maxID, nil
}

func (r *OrderRepository) GetProductByID(id int) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}