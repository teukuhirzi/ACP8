package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type OrderItem struct {
	ID        int `gorm:"size:36;not null;uniqueIndex;primary_key"`
	OrderID   int `gorm:"size:36;index"`
	ProductID int `gorm:"size:36;index"`
	Qty       int
	BasePrice decimal.Decimal `gorm:"type:decimal(16,2)"`
	TaxAmount decimal.Decimal `gorm:"type:decimal(16,2)"`
	SubTotal  decimal.Decimal `gorm:"type:decimal(16,2)"`
	Name      string          `gorm:"size:255"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
