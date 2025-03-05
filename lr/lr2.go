package lr

import (
	"github.com/cipher/des"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Lr2() {
	route := gin.Default()
	route.Use(CORSMiddleware())
	route.POST("/api/code", func(ctx *gin.Context) {
		type Request struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		}
		var request Request
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		code, err := des.Code(request.Text, request.Key)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"code": code})
	})
	route.POST("/api/decode", func(ctx *gin.Context) {
		type Request struct {
			Text string `json:"text"`
			Key  string `json:"key"`
		}
		var request Request
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		decode, err := des.Decode(request.Text, request.Key)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"decode": decode})
	})
	if err := route.Run(":8088"); err != nil {
		log.Fatal(err)
	}
}
