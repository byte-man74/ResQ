package user

import (
	"net/http"
	"resq/pkg/models"
	"resq/pkg/utils"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
}

type userController struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return &userController{service: service}
}

func (u *userController) CreateUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationErrors(err, user)})
		return
	}

	updatedUser, err := u.service.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": updatedUser})

}
