package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"resq/pkg/dto"
	"resq/pkg/models"
	"resq/pkg/utils"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	AuthorizeUser(ctx *gin.Context)
}

type userController struct {
	service UserService
}

const (
	requestError = "error"
	requestData  = "data"
)

func NewUserController(service UserService) UserController {
	return &userController{service: service}
}

func (u *userController) CreateUser(ctx *gin.Context) {

	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{requestError: utils.FormatValidationErrors(err, user)})
		return
	}

	updatedUser, err := u.service.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{requestData: updatedUser})
}
func (u *userController) AuthorizeUser(ctx *gin.Context) {
	var login dto.LoginRequestDTO

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{requestError: utils.FormatValidationErrors(err, login)})
		return
	}

	token, err := u.service.AuthorizeUser(&login)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{requestError: "authorization error"})
		return
	}

	if token == "" { 
		ctx.JSON(http.StatusInternalServerError, gin.H{requestError: "invalid token generated"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{requestData: token})
}
