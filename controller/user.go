package controller

import (
	"net/http"

	"app/middleware"
	"app/model"
	"app/request"
	"app/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
)

func LoginUser(c *gin.Context) {
	middleware.AuthMiddleware.LoginHandler(c)
}

// ユーザーを作成
func CreateUser(c *gin.Context) {
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

	// パスワードを暗号化
	password, err := util.PasswordEncrypt(req.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// userのデータを格納
	user := model.User{
		Name:        req.Name,
		Email:       req.Email,
		Password:    password,
		Description: req.Description,
	}

	// userをAPIに格納
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// パスワードを空にする。
	req.Password = ""

	c.JSON(http.StatusOK, req)
}
