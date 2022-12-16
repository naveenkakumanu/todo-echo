package routes

import (
	"github.com/labstack/echo"
	"github.com/naveenkakumanu/todo-echo/controllers"
)

func Routes() *echo.Echo {

	router := echo.New()
	// Create User
	router.POST("/user/create", controllers.CreateUser)
	router.GET("/user/read/:id", controllers.ReadUser)
	router.PUT("/user/update/:id", controllers.UpdateUser)
	router.DELETE("/user/delete/:id", controllers.DeleteUser)

	// Create Task
	router.POST("/app/create", controllers.Create)
	router.GET("/app/read/:id", controllers.Read)
	router.PUT("/app/update/:id", controllers.Update)
	router.DELETE("/app/delete/:id", controllers.Delete)
	return router
}
