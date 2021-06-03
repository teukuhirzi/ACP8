package database

import (
	"acp8/configs"
	"acp8/models"

	"github.com/labstack/echo"
)

func CreateItem(c echo.Context) (interface{}, interface{}) {
	item := models.Product{}
	c.Bind(&item)
	if err := configs.DB.Save(&item).Error; err != nil {
		return nil, err
	}
	configs.DB.Joins("ProductCategory").Where(&item).First(&item)

	return item, nil
}
