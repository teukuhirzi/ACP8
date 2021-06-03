package models

import (
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type ProductCategory struct {
	ID           uint           `gorm:"primaryKey" json:"item_category_id" form:"item_category_id"`
	CreatedAt    time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	CategoryName string         `gorm:"type:varchar(100);unique;not null" json:"category_name" form:"category_name"`
}

func (i *ProductCategory) First(c echo.Context, DB *gorm.DB) error {
	if err := DB.Where(i).First(i).Error; err != nil {
		return err
	}
	return nil
}
