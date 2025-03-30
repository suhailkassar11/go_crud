package controllers

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/suhailkassar11/go-crud/initializers"
	"github.com/suhailkassar11/go-crud/models"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = os.Getenv("JWT_SECRET")

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "please enter your data properly"})
		return
	}

	var existingUser models.User
	err = initializers.DB.Where("email=?", &user.Email).First(&existingUser).Error

	if err == nil {
		c.JSON(400, gin.H{"error": err, "message": "user already exists"})
		return
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "password hashing failed"})
	}

	user.Password = string(hasedPassword)

	err = initializers.DB.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "error in creating new user"})
		return
	}

	c.JSON(200, gin.H{"user": user, "message": "user created successfully"})

}

func LoginUser(c *gin.Context) {
	var loginUser models.UserLogin
	err := c.BindJSON(&loginUser)
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "please enter your credentials"})
		return
	}

	var user models.User
	err = initializers.DB.Where("email=?", loginUser.Email).First(&user).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "invalid email or password"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "hashing password is invalid"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(SecretKey))

	if err != nil {
		c.JSON(400, gin.H{"error": err, "token": tokenString, "message": "error in generating token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString, "message": "Login successfull"})
}

func FindAllUser(c *gin.Context) {
	var users []models.User
	err := initializers.DB.Find(&users).Error

	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "error in finding all users"})
		return
	}
	c.JSON(200, gin.H{"all users": users, "message": "successfully find all users"})
}

func FindOneUser(c *gin.Context) {

	id := c.Param("id")

	var userFound models.User
	err := initializers.DB.Where("id=?", id).First(&userFound).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "error in finding"})
		return
	}

	c.JSON(200, gin.H{"user": userFound, "message": "user found"})
}

func UpdateUser(c *gin.Context) {

	id := c.Param("id")

	var userFound models.User
	err := initializers.DB.Where("id=?", id).First(&userFound).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "user not exists"})
		return
	}

	initializers.DB.Where("id=?", id).Updates(&models.User{
		Username: c.PostForm("username"),
		Email:    c.PostForm("email"),
	})

	c.JSON(200, gin.H{"user": userFound, "message": "user found"})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := initializers.DB.Delete(&models.User{}, id).Error
	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "error in deleting user not found"})
		return
	}

	c.JSON(200, gin.H{"message": "user deleted successfully"})
}
