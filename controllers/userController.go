package controllers

import (
	"github.com/gin-gonic/gin"

	"vinicius-permor/apiGin/services"
)

type UserControllers struct {
	service *services.UserService
}

func NewUserController(services *services.UserService) *UserControllers {
	return &UserControllers{service: services}
}

func (usrControllers *UserControllers) ListAllUser(c *gin.Context) {
}
