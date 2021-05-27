package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID               string          `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Categories       []Category      `gorm:"many2many:product_categories;"`
	Name             string          `gorm:"size:255"`
	Price            decimal.Decimal `gorm:"type:decimal(16,2);"`
	Stock            int
	ShortDescription string `gorm:"type:text"`
	Description      string `gorm:"type:text"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
