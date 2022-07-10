package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"uj-ebiznes-go-crud/models"
)

func GetProducts(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var products []models.Product
		db.Find(&products)
		return c.JSON(http.StatusOK, products)
	}
}

func GetProduct(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var product models.Product
		db.Where("id = ?", id).Find(&product)
		return c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var product = new(models.Product)
		c.Bind(product)
		db.Create(&product)
		return c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var product = new(models.Product)
		db.Where("id = ?", id).Find(&product)
		c.Bind(product)
		db.Save(&product)
		return c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		db.Delete(&models.Product{}, id)
		return c.NoContent(http.StatusNoContent)
	}
}
