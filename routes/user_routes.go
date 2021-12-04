package routes

import (
	"github.com/Xsidelight/go_server/controllers"
	"github.com/gin-gonic/gin"
)

//UserRoutes - defining user routes
func UserRoutes(incomingRoutes *gin.Engine) {
    incomingRoutes.POST("/users/signup", controllers.SignUp())
    incomingRoutes.POST("/users/login", controllers.Login())
}
