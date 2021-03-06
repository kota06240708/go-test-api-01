package middleware

import (
	"fmt"
	"time"

	"app/model"
	"app/request"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/theckman/go-securerandom"
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

// リフレッシュトークン
var maxRefresh = time.Hour * 24 * 30 * 6

// ユーザー認証
func authMiddleware() *jwt.GinJWTMiddleware {
	var authMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
		Realm:      "test zone",
		Key:        []byte("secret key"),
		Timeout:    time.Hour * 24 * 30,
		MaxRefresh: maxRefresh,

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

			// 有効期限を発行
			now := time.Now()
			expire := now.Add(maxRefresh)

			// ランダムのバイト数を生成（リフレッシュトークンになる）
			rStr, err := securerandom.Base64OfBytes(64)

			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			// リフレッシュトークンのモデル
			refreshToken := &model.RefreshToken{
				Token:  rStr,
				Expire: expire,
			}

			// リフレッシュトークンをDBに格納
			if err := DB.Create(refreshToken).Error; err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			// データに格納
			c.Set("refreshToken", refreshToken)

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

		// ログイン時に返すres
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			// dbを取得
			refreshToken := c.MustGet("refreshToken").(*model.RefreshToken)

			c.JSON(code, gin.H{
				"code":         code,
				"token":        token,
				"refreshToken": refreshToken.Token,
				"expire":       expire.Format(time.RFC3339),
			})
		},

		// ===========================================
		// Middleware (token確認 +  tokenのデータを取得)
		// ===========================================

		// MiddlewareFuncを使うと呼ばれる
		// tokenの中身を確認、idを取得してユーザー情報があるか確認する
		IdentityHandler: func(c *gin.Context) interface{} {

			// dbを取得
			DB := c.MustGet("db").(*gorm.DB)

			// tokenの中身を確認
			claims := jwt.ExtractClaims(c)

			// ユーザーIDを取得
			id := claims["userID"].(float64)

			// dbのデータを格納するデータ
			user := &model.User{}

			fmt.Println(id)

			// データがあるか確認
			if err := DB.Where("ID = ?", id).First(user).Error; err != nil {
				fmt.Println("ばか")

				// 何もなかったら変換する
				return nil
			}

			// ログインしたユーザー情報を返す
			return user
		},

		// MiddlewareFuncを使うと呼ばれる。
		// IdentityHandlerで返された値が入る
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// ログインしたユーザー情報があるか確認
			if v, ok := data.(*model.User); ok {

				// ログインしたユーザー情報をginに格納
				c.Set("currentUser", v)
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
