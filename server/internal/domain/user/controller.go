package user

import (
	"net/http"
	"resq/pkg/constants"
	"resq/pkg/dto"
	"resq/pkg/models"
	"resq/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	AuthorizeUser(ctx *gin.Context)
	GetUserProfileInformation(ctx *gin.Context)
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
		ctx.JSON(http.StatusBadRequest, gin.H{constants.RequestError: utils.FormatValidationErrors(err, user)})
		return
	}

	createdUser, err := u.service.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{constants.RequestData: createdUser})
}
func (u *userController) AuthorizeUser(ctx *gin.Context) {
	var login dto.LoginRequestDTO

	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{constants.RequestError: utils.FormatValidationErrors(err, login)})
		return
	}

	token, err := u.service.AuthorizeUser(&login)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{constants.RequestError: "authorization error"})
		return
	}

	if token == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{constants.RequestError: "invalid token generated"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{constants.RequestData: token})
}


func (u *userController) GetUserProfileInformation(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{constants.RequestError: "unauthorized"})
		return
	}

	if userID == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{constants.RequestError: "invalid user id"})
		return
	}

	userIdStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{constants.RequestError: "internal server error"})
		return
	}

	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{constants.RequestError: "internal server error"})
		return
	}

	userId := uint(userIdInt)


	result, err := u.service.GetUserProfileInformation(userId)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{constants.RequestError: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{constants.RequestData: result})
}
