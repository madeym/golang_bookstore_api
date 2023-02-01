package models

import (
	"belajar-api/helper"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          int
	Image       string
	Name        string
	Email       string
	Phone       string
	DateOfBirth time.Time
	CityID      int
	ProvinceID  int
	Address     string
	Password    string
	LastLogin   time.Time
}

func init() {
	db.AutoMigrate(&User{})
}

func LoginUser(email string, password string) (User, error) {
	var existUser User
	err := db.First(&existUser, "email=?", email).Error
	// User Ketemu, Login
	if err == nil && existUser.Email != "" {
		// Check Password
		_, err := CheckPasswordHash(existUser.Password, password)
		existUser.LastLogin = time.Now()
		db.Save(&existUser)
		return existUser, err
	}

	// User Tidak Ketemu, Return Error
	return existUser, err
}

func RegisterUser(user User) (User, string, error) {
	var existUser User
	err := db.First(&existUser, "email=?", user.Email).Error

	if err != nil && existUser.Email == "" {
		p, err := HashPassword(user.Password)
		if err == nil {
			user.Password = p
		}
		err2 := db.Create(&user).Error
		return user, "", err2
	}

	if err == nil && existUser.Email != "" {
		return existUser, "User Already Registered", err
	}

	return existUser, "", err
}

func DetailUser() (User, error) {
	var user User
	err := db.First(&user, "ID = ?", helper.UserData.ID).Error

	return user, err
}

func ChangePasswordUser(newPassword string) (User, string, error) {
	var user User
	hashpassword, _ := HashPassword(newPassword)
	err := db.Model(&user).Where("id = ?", helper.UserData.ID).Updates(User{Password: hashpassword}).Error
	if err != nil {
		err = db.First(&user, helper.UserData.ID).Error
		return user, "", err
	}
	return user, "", err
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(hash string, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	if err != nil {
		return false, err
	}

	return true, nil
}

func GenerateUserToken(user User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["level"] = "application"
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}

func UpdateUser(user User) (User, error) {
	err := db.Model(&user).Where("id = ?", helper.UserData.ID).Updates(&user).Error

	return user, err
}
