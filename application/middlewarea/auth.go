/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  auth
 * @Time: 2020/9/17 10:01 上午
 */

package middlewarea

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Claims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Scope    uint8  `json:"scope"`
	jwt.StandardClaims
}

var jwtSecret = []byte("gsax")

func Authorize() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Token")
		if token == "" {
			context.JSON(http.StatusOK, gin.H{"errorCode": 40010, "message": "请求未携带用户token，无权限访问"})
			context.Abort()
			return
		}
		claims, err := ParseToken(token)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{"errorCode": 40016, "message": "登录令牌过期或无效"})
			context.Abort()
			return
		}
		context.Set("user_id", claims.ID)
		context.Set("username", claims.Username)
		context.Next()
	}
}

func GenerateToken(id string, username string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(720 * time.Hour) //3小时
	claims := Claims{
		ID:       id,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //token过期时间
			Issuer:    "zz",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
