package controllers

import (
	"belajar-api/helper"
	"belajar-api/models"
	"belajar-api/request"
	"belajar-api/response"
	"time"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	var userRequest request.UserLoginRequest

	err := c.ShouldBindJSON(&userRequest)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	user, e := models.LoginUser(userRequest.Email, userRequest.Password)
	if e != nil {
		helper.ResponseError(c, "", e)
		return
	}

	token, err := models.GenerateUserToken(user)

	response.UserResponse1(c, "User Login Successfully", token, user)
}

func RegisterUser(c *gin.Context) {
	var userRequest request.UserRegisterRequest

	err := c.ShouldBindJSON(&userRequest)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	newUser := models.User{
		Name:      userRequest.Name,
		Email:     userRequest.Email,
		Password:  userRequest.Password,
		LastLogin: time.Now(),
	}

	user, msg, e := models.RegisterUser(newUser)
	if msg != "" || err != nil {
		helper.ResponseError(c, msg, e)
		return
	}

	response.UserResponse1(c, "Successfully Registered New User", "", user)
}

func DetailUser(c *gin.Context) {
	user, e := models.DetailUser()
	if e != nil {
		helper.ResponseError(c, "", e)
		return
	}

	response.UserResponse1(c, "User Data Successfully Displayed", "", user)
}

func ChangePasswordUser(c *gin.Context) {
	var userRequest request.UserChangePasswordRequest

	err := c.ShouldBindJSON(&userRequest)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	user, msg, err := models.ChangePasswordUser(userRequest.Password)
	if msg != "" || err != nil {
		helper.ResponseError(c, msg, err)
		return
	}

	response.UserResponse1(c, "Successfully Change Password User", "", user)
}

func UpdateUser(c *gin.Context) {
	var userRequest request.UpdateUserRequest

	err := c.ShouldBindJSON(&userRequest)
	if helper.ShouldBindJSONError(c, err) {
		return
	}

	updateUser := models.User{
		Image:       userRequest.Image,
		Name:        userRequest.Name,
		Email:       userRequest.Email,
		Phone:       userRequest.Phone,
		DateOfBirth: userRequest.DateOfBirth,
		Address:     userRequest.Address,
	}

	user, err := models.UpdateUser(updateUser)
	if err != nil {
		helper.ResponseError(c, "", err)
		return
	}

	response.UserResponse1(c, "Successfully Update User", "", user)
}
