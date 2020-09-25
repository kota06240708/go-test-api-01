package controller

import (
	"net/http"

	"app/model"
	"app/request"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

// getの処理
func GetUsers(c *gin.Context) {
	var users []model.User

	DB := c.MustGet("db").(*gorm.DB)

	DB.Find(&users)

	//
	c.JSON(http.StatusOK, users)
}

// postの処理
func PostUser(c *gin.Context) {
	var req request.User
	DB := c.MustGet("db").(*gorm.DB)

	validate := validator.New()

	// reqのjsonデータを取得
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate
	if err := validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userのデータを格納
	user := model.User{
		Name:    req.Name,
		Context: req.Context,
	}

	// userをAPIに格納
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
	})
}
