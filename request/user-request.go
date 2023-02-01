package request

import "time"

type UserRegisterRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserChangePasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Name        string    `json:"name" binding:"required"`
	Image       string    `json:"image"`
	Email       string    `json:"email" binding:"required,email"`
	Phone       string    `json:"phone" binding:"required"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Address     string    `json:"address" binding:"required"`
}
