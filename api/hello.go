package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// vercel Function,必须有一个导出函数如下,才能部署到Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}

var router *gin.Engine

func init() {
	router = gin.Default()
	RegisterHandler(router)
}

// 支持embed file,简单的页面可以使用embed,只需要一个可执行文件
// 注意embed目录需要放到api下与接口同级!
// var files embed.FS

func RegisterHandler(r *gin.Engine) {
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"code": 0,
			"msg":  "hello world",
		})
	})
}
