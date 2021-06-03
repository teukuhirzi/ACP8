package controllers

import (
	"acp8/libs/database"
	"acp8/middlewares"
	"acp8/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetUserControllers(c echo.Context) error {

	data, err := database.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": "Failed get data User",
			"data":    nil,
		})
	}

	// userId
	userId := middlewares.ExtractJWTToUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"userId":  userId,
		"status":  "success",
		"message": "Success get data user",
		"data":    data,
	})
}

func RegisterControllers(c echo.Context) error {
	var userRegister models.UserRegister
	c.Bind(&userRegister)

	data, err := database.RegisterUser(userRegister)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": "Failed insert data User",
			"data":    nil,
		})
	}

	dataUserDB := data.(models.User)
	token, err := middlewares.GenerateJWT(int(dataUserDB.ID))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    500,
			"status":  "error",
			"message": "Failed generate token",
			"data":    nil,
		})
	}

	userResponse := models.UserResponse{
		dataUserDB.ID,
		dataUserDB.CreatedAt,
		dataUserDB.UpdatedAt,
		dataUserDB.Name,
		dataUserDB.Email,
		token,
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"status":  "success",
		"message": "Success insert data user",
		"data":    userResponse,
	})
}
