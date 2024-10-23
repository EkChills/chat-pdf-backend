package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"

	"path/filepath"
)

func UploadPdfs(context *gin.Context) {
	file, err := context.FormFile("file")

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	
if filepath.Ext(file.Filename) != ".pdf" {
	context.JSON(http.StatusBadRequest, map[string]string {
		"message": "Invalid file type: file must be a pdf",
	})
	return
}

tempDir, err := os.MkdirTemp("", "pdf-uploads-*")

if err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{
		"message": "could not create temp dir",
	})

	return
}

filePath := filepath.Join(tempDir, file.Filename)

if err := context.SaveUploadedFile(file, filePath); err != nil {
	context.JSON(http.StatusInternalServerError, gin.H{
		"message": "Error uploading file",
	})

	return
}

seekFile, err := os.Open(filePath)




}

func extractTextFromPdf(filePath string, fileSeeker io.ReadSeeker, outDir string) (string, error) {

	// cfg := pdfcpu.DefaultPageConfiguration()
	// outPutFile := filepath.Join(outDir, "extracted-texts")
	// api.ExtractContent(fileSeeker, outPutFile, "extracted-text", []string{}, cfg)
}
