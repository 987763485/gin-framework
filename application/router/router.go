/**
 * @Author: wuchunle<wuchunle@gsaxns.com>
 * @Version: 1.0.0
 * @Description:
 * @File:  router
 * @Time: 2020/9/17 9:47 上午
 */

package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var APP = gin.Default()

func init() {
	APP.Any("/", func(context *gin.Context) {
		context.Redirect(http.StatusMovedPermanently, "www")
	})
	APP.StaticFS("www", http.Dir("www"))
	APP.StaticFS("static", http.Dir("static"))
}
