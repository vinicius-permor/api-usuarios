package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"vinicius-permor/apiGin/src/models"
	"vinicius-permor/apiGin/src/services"
)

type UserControllers struct {
	service     *services.UserService
	authService *services.AuthService
}

func NewUserController(services *services.UserService, authService *services.AuthService) *UserControllers {
	return &UserControllers{
		service:     services,
		authService: authService,
	}
}

func (usrControllers *UserControllers) Login(c *gin.Context) {
	var usersCredentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&usersCredentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	users, err := usrControllers.authService.Login(usersCredentials.Email, usersCredentials.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "login criado com sucesso",
		"data":    users,
	})
}

func (usrControllers *UserControllers) CreateUser(c *gin.Context) {
	var users models.Users
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	if err := users.Prepare("create"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}
	id, err := usrControllers.service.CreateUser(&users)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}

	users.ID = int(id)
	users.Password = "" // nao vai retornar a senha

	c.JSON(http.StatusCreated, gin.H{
		"message": "cliente criado com sucesso",
		"data":    users,
	})
}

func (usrControllers *UserControllers) ListAllUser(c *gin.Context) {
	users, err := usrControllers.service.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (usrControllers *UserControllers) SearchUser(c *gin.Context) {
	id := c.Param("id")
	user, err := usrControllers.service.SearchUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}
	user.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (usrControllers *UserControllers) UpadateUser(c *gin.Context) {
	var users models.Users
	id := c.Param("id")
	if err := c.ShouldBindJSON(&users); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}
	if err := users.Prepare("update"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}
	err := usrControllers.service.UpdateUserID(id, &users)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"erro": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "usuario autalizado com sucesso",
		"data":    id,
	})
}

func (usrControllers *UserControllers) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := usrControllers.service.DeleteUserID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"erro": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "cliente foi deletado",
	})
}
