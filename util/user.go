package util

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// requestを取得
func GetRequest(c *gin.Context, data interface{}) {
	validate := validator.New()

	// reqのjsonデータを取得
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(nil)
	}

	// validate
	if err := validate.Struct(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		panic(nil)
	}

	fmt.Println("通ってします")
}
