package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"resq.com/resq/server/dto"
	"resq.com/resq/server/internal/service"
	"resq.com/resq/server/internal/utils"
	"resq.com/resq/server/models"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	UpdateUserBasicInformation(ctx *gin.Context)
	UpdateUserSocialInformation(ctx *gin.Context)
	GetFullUserInformation(ctx *gin.Context)
	AuthenticateUser(ctx *gin.Context)
}




type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{service: service}
}

func (c *userController) CreateUser(ctx *gin.Context) {
	var request dto.ICreateAccount
	if err := ctx.ShouldBindJSON(&request); err != nil {
		
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.FormatValidationErrors(err)})
		return
	}

	// newUser, err := c.service.CreateUser(&request.User, &request.RegistrationLocationInformation, &request.Location)

	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	ctx.JSON(http.StatusCreated, gin.H{"data": "pass"})
}

func (c *userController) UpdateUserBasicInformation(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := c.service.UpdateUserBasicInformation(&user)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedUser})
}

func (c *userController) UpdateUserSocialInformation(ctx *gin.Context) {
	var socialInformation models.UserSocialInformation

	if err := ctx.ShouldBindJSON(&socialInformation); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSocialInformation, err := c.service.UpdateUserSocialInformation(&socialInformation)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedSocialInformation})
}

func (c *userController) GetFullUserInformation(ctx *gin.Context) {

	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}
	fullInformation, err := c.service.GetFullUserInformation(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": fullInformation})

}


func (c *userController) AuthenticateUser(ctx *gin.Context) {
	var login dto.ILogin

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authenticatedPayload, err := c.service.AuthenticateUser(&login)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if authenticatedPayload == nil || authenticatedPayload.User == nil {
		ctx.JSON(http.StatusUnauthorized , gin.H{"error": "authentication failed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": authenticatedPayload})
}
