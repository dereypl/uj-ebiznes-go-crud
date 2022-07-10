package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"uj-ebiznes-go-crud/models"
)

func GetUsers(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var users []models.User
		db.Find(&users)
		return c.JSON(http.StatusOK, users)
	}
}

func GetUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var user models.User
		db.Where("id = ?", id).Preload("Orders").Preload("CreditCard").Find(&user)
		return c.JSON(http.StatusOK, user)
	}
}

func CreateUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var user = new(models.User)
		c.Bind(user)
		db.Create(&user)
		return c.JSON(http.StatusOK, user)
	}
}

func UpdateUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var user = new(models.User)
		db.Where("id = ?", id).Find(&user)
		c.Bind(user)
		db.Save(&user)
		return c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		db.Delete(&models.User{}, id)
		return c.NoContent(http.StatusNoContent)
	}
}
