package domain

type Order struct {
	ID	 int		`gorm:"primaryKey"`
	UserID int		`gorm:"primaryKey"`
	ProductID int	`gorm:"primaryKey"`
	Quantity int	`gorm:"not null"`
}