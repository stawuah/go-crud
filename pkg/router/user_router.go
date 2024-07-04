package router

import (
	handler "github/stawuah/go-crud/internal/modules/users/handlers"
	service "github/stawuah/go-crud/internal/modules/users/services"

	"github.com/gin-gonic/gin"
)

// NewRouter sets up the Gin router and routes, injecting the necessary service
func NewRouter(userService service.UserService) *gin.Engine {
	r := gin.Default()

	userHandler := handler.NewUserHandler(userService)

	r.GET("/users/:id", userHandler.GetUser)
	r.POST("/users", userHandler.RegisterUser)

	return r
}
