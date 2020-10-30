package controller

import (
	"fmt"
	"net/http"

	"app/model"
	"app/request"
	"app/util"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm"
)

// コメントを追加
func CreateComment(c *gin.Context) {
	var req request.Comment
	DB := c.MustGet("db").(*gorm.DB)

	session := sessions.Default(c)

	session.Set("UserId", "1211111")
	session.Set("Test", "ばか")
	session.Set("tqnnmnk", "ばか")
	session.Save()

	fmt.Println(session.Get("tqnnmnk"))

	currentUser := c.MustGet("currentUser").(*model.User)

	// パラメーターを取得
	util.GetRequest(c, &req)

	// userのデータを格納
	comment := &model.Comment{
		UserId:  currentUser.ID,
		Comment: req.Comment,
	}

	getComments := &[]model.Comment{}

	user := &model.User{}

	// コメントをAPIに保存
	if err := DB.Create(comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Preload("Comment").Preload("Comment.User").Where("ID = ?", currentUser.ID).First(user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int

	if err := DB.Preload("User").Find(&getComments).Count(&count).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}
