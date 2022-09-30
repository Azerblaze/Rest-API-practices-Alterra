package route

import (
	"praktikum/constants"
	"praktikum/controller"
	m "praktikum/middleware"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route / to handler function
	e.POST("/login", controller.LoginUsersController)

	config := middleware.JWTConfig{
		Claims:     &m.JwtCustomClaims{},
		SigningKey: []byte(constants.SECRET_JWT),
	}

	//metode group tidak dapat digunakan
	//sudah dicoba tapi ketika dijalankan, tidak dilakukan autentikasi
	e.GET("/users", controller.GetUsersController, middleware.JWTWithConfig(config))
	e.GET("/users/:id", controller.GetUserController, middleware.JWTWithConfig(config))
	e.POST("/users", controller.CreateUserController)
	e.DELETE("/users/:id", controller.DeleteUserController, middleware.JWTWithConfig(config))
	e.PUT("/users/:id", controller.UpdateUserController, middleware.JWTWithConfig(config))

	e.GET("/books", controller.GetBooksController)
	e.GET("/books/:id", controller.GetBookController)
	e.POST("/books", controller.CreateBookController, middleware.JWTWithConfig(config))
	e.DELETE("/books/:id", controller.DeleteBookController, middleware.JWTWithConfig(config))
	e.PUT("/books/:id", controller.UpdateBookController, middleware.JWTWithConfig(config))

	return e
}
