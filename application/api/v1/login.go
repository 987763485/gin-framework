/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  login
 * @Time: 2020/9/17 10:01 上午
 */

package v1

import (
	"github.com/987763485/gin-framework/application/middlewarea"
	"github.com/987763485/gin-framework/application/models"
	"github.com/987763485/gin-framework/application/pkg/orm"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login() gin.HandlerFunc {
	type UserLoginData struct {
		LoginName string `form:"login_name" json:"login_name" binding:"required"`
		Password  string `form:"password" json:"password" binding:"required"`
	}
	return func(context *gin.Context) {
		var params UserLoginData
		if err := context.ShouldBind(&params); err != nil {
			context.JSON(http.StatusOK, gin.H{"errorCode": 10012, "message": err.Error()})
			context.Abort()
			return
		}
		user := models.User{}
		if isNotFound := orm.SlaveDB().Where("login_name=?", params.LoginName).First(&user).RecordNotFound(); isNotFound {
			context.JSON(http.StatusOK, gin.H{"errorCode": 10002, "message": "用户名不存在"})
			context.Abort()
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(params.Password)); err != nil {
			context.JSON(http.StatusOK, gin.H{"errorCode": 10003, "message": "密码错误"})
			context.Abort()
			return
		}
		if token, err := middlewarea.GenerateToken(user.ID, user.Name); err != nil {
			context.JSON(http.StatusOK, gin.H{"errorCode": 10003, "message": err.Error()})
			context.Abort()
			return
		} else {
			context.JSON(http.StatusOK, gin.H{"code": 200, "message": "login success", "data": gin.H{"token": token}})
			context.Abort()
			return
		}
	}
}
