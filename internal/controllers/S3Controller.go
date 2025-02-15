package controllers

import (
	"STDE_proj/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
)

func PostFileHandler(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Необходимо загрузить файл"})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось открыть файл"})
		return
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	fileName := fileHeader.Filename

	url, err := services.PostFile(context.TODO(), file, fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func DeleteFileHandler(c *gin.Context) {
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя файла не указано"})
		return
	}
	err := services.DeleteFile(context.TODO(), fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Файл удалён"})
}

func GetFileURLHandler(c *gin.Context) {
	fileName := c.Param("filename")
	bucketName := c.Param("bucket")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя файла не указано"})
		return
	}
	url := services.GetFileURL(context.TODO(), bucketName+"/"+fileName)

	c.JSON(http.StatusOK, gin.H{"url": url})
}

func DownloadFileHandler(c *gin.Context) {
	fileName := c.Param("filename")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя файла не указано"})
		return
	}
	data, err := services.DownloadFile(context.TODO(), fileName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Отдаем данные как бинарный поток
	c.Data(http.StatusOK, "application/octet-stream", data)
}

func ListFilesHandler(c *gin.Context) {
	prefix := c.Query("prefix") // если не указан, будет пустая строка
	res, err := services.ListFiles(context.TODO(), prefix)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"files": res.Contents})
}

func ListBucketsHandler(c *gin.Context) {
	res, err := services.ListBuckets(context.TODO())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"buckets": res.Buckets})
}
