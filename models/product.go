package models

import (
	"time"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type Product struct {
	ID                uint            `gorm:"primaryKey" json:"item_id" form:"item_id"`
	CreatedAt         time.Time       `json:"created_at" form:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at" form:"updated_at"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name              string          `gorm:"type:varchar(100);unique;not null" json:"name" form:"name"`
	Description       string          `gorm:"type:varchar(500)" json:"description" form:"description"`
	Stock             uint            `json:"stock" form:"stock"`
	Price             uint            `json:"price" form:"price"`
	ProductCategoryID uint            `gorm:"not null" json:"item_category_id" form:"item_category_id"`
	ProductCategory   ProductCategory `json:"item_category" form:"item_category"`
}

//create in database
func (i *Product) Create(c echo.Context, DB *gorm.DB) error {
	c.Bind(i)
	return DB.Save(&i).Error
}

//find in the database
func (i *Product) Find(c echo.Context, DB *gorm.DB) error {
	c.Bind(i)
	return DB.Where(i).First(i).Error
}

type ProductResponse struct {
	Code    uint    `json:"code" form:"code"`
	Status  string  `json:"status" form:"status"`
	Message string  `json:"message" form:"message"`
	Data    Product `json:"data,omitempty" form:"data"`
}

type ProductResponseArr struct {
	Code    uint      `json:"code" form:"code"`
	Status  string    `json:"status" form:"status"`
	Message string    `json:"message" form:"message"`
	Data    []Product `json:"data,omitempty" form:"data"`
}
