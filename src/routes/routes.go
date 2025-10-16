package routes

import (
	"database/sql"

	"vinicius-permor/apiGin/src/controllers"
	"vinicius-permor/apiGin/src/repositories"
	"vinicius-permor/apiGin/src/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, db *sql.DB) {
	userRepo := repositories.NewUserRepository(db)
	userServ := services.NewUserService(userRepo)
	userAuth := services.NewAuthService(userRepo)
	userCrontrol := controllers.NewUserController(userServ, userAuth)

	users := r.Group("/users")
	users.GET("", userCrontrol.ListAllUser)
	users.GET("/:id", userCrontrol.SearchUser)
	users.POST("", userCrontrol.CreateUser)
	users.PUT("/:id", userCrontrol.UpadateUser)
	users.DELETE("/:id", userCrontrol.DeleteUser)
}
