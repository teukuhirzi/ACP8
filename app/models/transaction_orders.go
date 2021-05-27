package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction_Order struct {
	ID         string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	UserID     string `gorm:"size:36;index"`
	User       User
	CartIds    []int
	GrandTotal decimal.Decimal `gorm:"type:decimal(16,2)"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
