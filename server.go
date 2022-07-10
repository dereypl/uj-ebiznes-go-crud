package main

import (
	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"uj-ebiznes-go-crud/controllers"
	"uj-ebiznes-go-crud/models"
)

func main() {

	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.CreditCard{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Category{})
	db.AutoMigrate(&models.Order{})

	CreditCard1 := models.CreditCard{Number: "1234567891", UserID: 1}
	CreditCard2 := models.CreditCard{Number: "1234567892", UserID: 2}
	CreditCard3 := models.CreditCard{Number: "1234567893", UserID: 3}
	db.Create(&CreditCard1)
	db.Create(&CreditCard2)
	db.Create(&CreditCard3)

	User1 := models.User{Email: "test1@test.uj", FirstName: "Test", LastName: "User1", CreditCard: CreditCard1}
	db.Create(&User1)

	db.Create(&models.User{Email: "test2@test.uj", FirstName: "Test", LastName: "User2", CreditCard: CreditCard2})
	db.Create(&models.User{Email: "test3@test.uj", FirstName: "Test", LastName: "User3", CreditCard: CreditCard3})

	Product1 := models.Product{Name: "Macbook Pro 1", Price: 12999.00, CategoryID: 1}
	Product2 := models.Product{Name: "Macbook Pro 2", Price: 12999.00, CategoryID: 1}
	Product3 := models.Product{Name: "Macbook Pro 3", Price: 12999.00, CategoryID: 1}
	db.Create(&Product1)
	db.Create(&Product2)
	db.Create(&Product3)

	Category1 := models.Category{Name: "Laptopy i komputery", Products: []models.Product{Product1, Product2, Product3}}
	db.Create(&Category1)

	db.Create(&models.Order{UserID: 1, Products: []models.Product{Product3}})

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/categories", controllers.GetCategories(db))
	e.GET("/categories/:id", controllers.GetCategory(db))
	e.POST("/categories", controllers.CreateCategory(db))
	e.PUT("/categories/:id", controllers.UpdateCategory(db))
	e.DELETE("/categories/:id", controllers.DeleteCategory(db))

	e.GET("/users", controllers.GetUsers(db))
	e.GET("/users/:id", controllers.GetUser(db))
	e.POST("/users", controllers.CreateUser(db))
	e.PUT("/users/:id", controllers.UpdateUser(db))
	e.DELETE("/users/:id", controllers.DeleteUser(db))

	e.GET("/products", controllers.GetProducts(db))
	e.GET("/products/:id", controllers.GetProduct(db))
	e.POST("/products", controllers.CreateProduct(db))
	e.PUT("/products/:id", controllers.UpdateProduct(db))
	e.DELETE("/products/:id", controllers.DeleteProduct(db))

	e.GET("/orders", controllers.GetOrders(db))
	e.GET("/orders/:id", controllers.GetOrder(db))
	e.POST("/orders", controllers.CreateOrder(db))
	e.PUT("/orders/:id", controllers.UpdateOrder(db))
	e.DELETE("/orders/:id", controllers.DeleteOrder(db))

	e.GET("/creditCards", controllers.GetCreditCards(db))
	e.GET("/creditCards/:id", controllers.GetCreditCard(db))
	e.POST("/creditCards", controllers.CreateCreditCard(db))
	e.PUT("/creditCards/:id", controllers.UpdateCreditCard(db))
	e.DELETE("/creditCards/:id", controllers.DeleteCreditCard(db))

	e.Logger.Fatal(e.Start(":1323"))
}
