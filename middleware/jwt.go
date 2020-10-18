package middleware

import (
	"fmt"
	"time"

	"app/model"
	"app/request"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var identityKey = "id"

// type login struct {
// 	Username string `form:"username" json:"username" binding:"required"`
// 	Password string `form:"password" json:"password" binding:"required"`
// }

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

// ユーザー認証
func authMiddleware() *jwt.GinJWTMiddleware {
	var authMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour * 24 * 30,
		MaxRefresh: time.Hour * 24 * 30 * 6,

		// ===========================================
		// ログイン時
		// ===========================================

		// ログイン時呼ばれる関数
		// 一番はじめにここに入る
		Authenticator: func(c *gin.Context) (interface{}, error) {
			// dbを取得
			DB := c.MustGet("db").(*gorm.DB)

			var loginVals request.Login
			if err := c.ShouldBindJSON(&loginVals); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			// apiからきたemailを取得
			email := loginVals.Email

			// dbのデータを格納するデータ
			user := &model.User{}

			// メールアドレスでユーザーを絞り込む
			// 無い場合はエラー
			if err := DB.Where("email = ?", email).First(user).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			// 取得したユーザーを返す
			return user, nil
		},

		// ログイン時に呼ばれる関数
		// tokenにデータを詰め込む
		// Authenticator
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				claims := jwt.MapClaims{
					"userID": v.ID,
					"name":   v.Name,
				}

				return claims
			}
			return jwt.MapClaims{}
		},

		// MiddlewareFuncを使うと呼ばれる
		// tokenの中身を確認、idを取得してユーザー情報があるか確認する
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["userID"].(string),
			}
		},

		// MiddlewareFuncを使うと呼ばれる。
		// IdentityHandlerで返された値が入る
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},

		// ===========================================
		// エラー時に呼ばれる
		// ===========================================

		// エラーした時に呼ばれる
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		// ログイン時に返すres
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"code":   code,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	return authMiddleware
}

// 継承
var AuthMiddleware = authMiddleware()

// func CreateToken(user request.User) (string, error) {

// 	// headerのセット
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	// claimsのセット
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["admin"] = true

// 	// ユーザーの一意識別子
// 	claims["name"] = user.Name
// 	claims["email"] = user.Email
// 	claims["iat"] = time.Now()
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

// 	// 電子署名
// 	tokenString, err := token.SignedString([]byte("secret"))

// 	fmt.Println("-----------------------------")
// 	fmt.Println("tokenString:", tokenString)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return tokenString, nil
// }

// var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
// 	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("SIGNINGKEY")), nil
// 	},
// 	SigningMethod: jwt.SigningMethodHS256,
// })
