package controller

import (
	"net/http"

	"example.com/resq/server/internal/models"
	"example.com/resq/server/internal/service"
	"github.com/gin-gonic/gin"
)


type UserController interface {
	CreateUser(ctx *gin.Context);
	UpdateUserBasicInformation(ctx *gin.Context);
	UpdateUserSocialInformation(ctx *gin.Context);
	GetFullUserInformation(ctx *gin.Context)
}


type userController struct {
	service service.UserService
}


func NewUserController (service service.UserService) UserController {
	return &userController{service: service}
}


func (c *userController) CreateUser(ctx *gin.Context) {
	var user models.User;
	var registrationLocationInformation models.UserRegistrationLocation;
	var location models.Location

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	newUser, err := c.service.CreateUser(&user, &registrationLocationInformation, &location)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()});
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newUser})
}


func (c *userController) UpdateUserBasicInformation (ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	updatedUser, err := c.service.UpdateUserBasicInformation(&user);

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedUser})
}


func (c *userController) UpdateUserSocialInformation (ctx *gin.Context) {
	var socialInformation models.UserSocialInformation;

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

	id := ctx.Param("id");


	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}
	fullInformation, err := c.service.GetFullUserInformation(id);

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": fullInformation})

}
