package routes

import (
	controller "golang-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incominRoutes *gin.Engine) {
	incominRoutes.GET("/users", controller.GetUsers())
	incominRoutes.GET("/users/:user_id", controller.GetUser())
	incominRoutes.POST("user/signup", controller.SignUp())
	incominRoutes.POST("user/login", controller.Login())

}
