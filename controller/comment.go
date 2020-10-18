package controller

import (
	"net/http"

	"app/request"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

// ユーザーを作成
func CreateComment(c *gin.Context) {
	var req request.Comment
	// DB := c.MustGet("db").(*gorm.DB)

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
	// comment := model.Comment{
	// 	UserId:  req.UserId,
	// 	Comment: req.Comment,
	// }

	// userをAPIに格納
	// if err := DB.Create(&comment).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": req.Comment,
	})
}
