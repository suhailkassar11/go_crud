package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/suhailkassar11/go-crud/initializers"
	"github.com/suhailkassar11/go-crud/models"
)

func CreateUser(c *gin.Context) {

	var user models.User

	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "please enter your data properly"})
		return
	}
	// if &user.Username == nil {
	// 	c.JSON(400, gin.H{"message": "please provide a username"})
	// } else if &user.Email == nil {
	// 	c.JSON(400, gin.H{"message": "please provide a email"})
	// } else {
	// 	c.JSON(400, gin.H{"message": "please provide a password"})
	// }
	var existingUser models.User
	err = initializers.DB.Where("email=?", &user.Email).First(&existingUser).Error

	if err == nil {
		c.JSON(400, gin.H{"error": err, "message": "user already exists"})
		return
	}

	err = initializers.DB.Create(&user).Error

	if err != nil {
		c.JSON(400, gin.H{"error": err, "message": "error in creating new user"})
		return
	}

	c.JSON(200, gin.H{"user": user, "message": "user created successfully"})

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
