package lr

import (
	"github.com/cipher/rsa"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Lr3() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	router.GET("/api/generate-keys", func(ctx *gin.Context) {
		enKey, deKey, err := rsa.GenerateKeys(1024)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"publicKey":  enKey,
			"privateKey": deKey,
		})
	})
	router.POST("/api/encrypt", func(ctx *gin.Context) {
		type Request struct {
			Text      string `json:"text"`
			PublicKey string `json:"publicKey"`
		}
		var request Request
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := rsa.Encrypt(request.Text, request.PublicKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"encryptedText": result,
		})
	})
	router.POST("/api/decrypt", func(ctx *gin.Context) {
		type Request struct {
			Text       string `json:"text"`
			PrivateKey string `json:"privateKey"`
		}
		var request Request
		if err := ctx.ShouldBind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		result, err := rsa.Decrypt(request.Text, request.PrivateKey)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"decryptedText": result,
		})
	})
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
