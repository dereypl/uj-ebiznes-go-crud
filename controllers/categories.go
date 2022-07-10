package controllers

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"uj-ebiznes-go-crud/models"
)

func GetCategories(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var categories []models.Category
		db.Find(&categories)
		return c.JSON(http.StatusOK, categories)
	}
}

func GetCategory(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var category models.Category
		db.Where("id = ?", id).Preload("Products").Find(&category)
		return c.JSON(http.StatusOK, category)
	}
}

func CreateCategory(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		var category = new(models.Category)
		c.Bind(category)
		db.Create(&category)
		return c.JSON(http.StatusOK, category)
	}
}

func UpdateCategory(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		var category = new(models.Category)
		db.Where("id = ?", id).Find(&category)
		c.Bind(category)
		db.Save(&category)
		return c.JSON(http.StatusOK, category)
	}
}

func DeleteCategory(db *gorm.DB) func(echo.Context) error {
	return func(c echo.Context) error {
		id := c.Param("id")
		db.Delete(&models.Category{}, id)
		return c.NoContent(http.StatusNoContent)
	}
}
