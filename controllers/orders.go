package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"uj-ebiznes-go-crud/models"
)

func GetOrders(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var orders []models.Order
		db.Find(&orders)
		return c.JSON(http.StatusOK, orders)
	}
}

func GetOrder(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var order models.Order
		db.Where("id = ?", id).Preload("Products").Find(&order)
		return c.JSON(http.StatusOK, order)
	}
}

func CreateOrder(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var order = new(models.Order)
		c.Bind(order)
		db.Create(&order)
		return c.JSON(http.StatusOK, order)
	}
}

func UpdateOrder(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var order = new(models.Order)
		db.Where("id = ?", id).Find(&order)
		c.Bind(order)
		db.Save(&order)
		return c.JSON(http.StatusOK, order)
	}
}

func DeleteOrder(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		db.Delete(&models.Order{}, id)
		return c.NoContent(http.StatusNoContent)
	}
}
