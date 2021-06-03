package controllers

import (
	"acp8/configs"
	"acp8/libs/database"
	"acp8/middlewares"
	"acp8/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetProductWIthParamsController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	model := models.Product{}

	if id > 0 {
		model.ID = uint(id)
	}
	products := []models.Product{}

	if err := configs.DB.Preload("ProductCategory").Where(&model).Find(&products).Error; err != nil {
		return c.JSON(http.StatusBadRequest, models.ProductResponseArr{
			Code:    http.StatusBadRequest,
			Status:  "failed",
			Message: "failed getting products",
		})
	}

	productResponseArr := models.ProductResponseArr{
		Code:    200,
		Status:  "success",
		Message: "success getting products",
		Data:    products,
	}
	return c.JSON(http.StatusOK, productResponseArr)

}

func RegisterProductController(c echo.Context) error {
	role := middlewares.ExtractTokenUserRole(c)
	if role == "admin" {
		item, err := database.CreateItem(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"status":  "failed",
				"message": "bad request",
				"data":    "",
			})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    http.StatusCreated,
			"status":  "success",
			"message": "product created",
			"data":    item,
		})

	}
	return c.JSON(http.StatusUnauthorized, map[string]interface{}{
		"code":    http.StatusUnauthorized,
		"status":  "failed",
		"message": "you have no permission",
		"data":    "",
	})
}
