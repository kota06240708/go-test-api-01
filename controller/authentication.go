package controller

import (
	"fmt"
	"net/http"

	"app/middleware"
	"app/model"
	"app/request"
	"app/util"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RefreshToken(c *gin.Context) {
	// 型を定義
	var req = request.RefreshToken{}

	// DBを定義
	DB := c.MustGet("db").(*gorm.DB)

	// reqを取得
	util.GetRequest(c, &req)

	// RefreshTokenのモデルを定義
	refreshToken := &model.RefreshToken{}

	if err := DB.First(refreshToken).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"RefreshToken": gin.H{"text": "refreshToken does not match", "tag": "notmatch"}})
		return
	}

	fmt.Println(refreshToken)

	// トークンを再発行
	middleware.AuthMiddleware.RefreshHandler(c)
}
