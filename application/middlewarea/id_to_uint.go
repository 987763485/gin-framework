/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  id_to_uint
 * @Time: 2020/9/17 10:00 上午
 */

package middlewarea

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IDToUint() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Param("id") != "" {
			paramID, err := strconv.ParseUint(context.Param("id"), 10, 32)
			if err != nil {
				context.JSON(http.StatusOK, gin.H{"errorCode": 10010, "message": "非法的id参数"})
				context.Abort()
				return
			}
			context.Set("id", uint(paramID))
		}
		context.Next()
	}
}
