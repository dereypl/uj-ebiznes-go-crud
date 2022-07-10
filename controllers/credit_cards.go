package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"uj-ebiznes-go-crud/models"
)

func GetCreditCards(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var creditCards []models.CreditCard
		db.Find(&creditCards)
		return c.JSON(http.StatusOK, creditCards)
	}
}

func GetCreditCard(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var creditCards models.CreditCard
		db.Where("id = ?", id).Preload("Products").Find(&creditCards)
		return c.JSON(http.StatusOK, creditCards)
	}
}

func CreateCreditCard(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var creditCards = new(models.CreditCard)
		c.Bind(creditCards)
		db.Create(&creditCards)
		return c.JSON(http.StatusOK, creditCards)
	}
}

func UpdateCreditCard(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var creditCards = new(models.CreditCard)
		db.Where("id = ?", id).Find(&creditCards)
		c.Bind(creditCards)
		db.Save(&creditCards)
		return c.JSON(http.StatusOK, creditCards)
	}
}

func DeleteCreditCard(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		db.Delete(&models.CreditCard{}, id)
		return c.NoContent(http.StatusNoContent)
	}
}
