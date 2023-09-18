package routes

import (
	"gorest-10/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, ini tugas rest api golang dengan tema stok item. Tambahkan /items pada url")
	})
	e.GET("/items", controllers.GetAllItems)
	e.GET("/items/:id", controllers.GetItemById)
	e.POST("/items", controllers.CreateItem)
	e.PUT("items/:id", controllers.UpdateItem)
	e.DELETE("items/:id", controllers.DeleteItemById)
}
