package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadPdfs(context *gin.Context) {
	file, err := context.FormFile("file")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
	}

	_ , err = file.Open()

}
