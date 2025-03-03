package controllers

import (
	"STDE_proj/internal/services"
	"context"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
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
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось прочитать файл"})
		return
	}

	objectKey := strings.TrimLeft(c.Param("key"), "/")
	if objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Не указан ключ объекта"})
		return
	}

	url, err := services.PostFile(context.TODO(), fileBytes, objectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func DeleteFileHandler(c *gin.Context) {
	objectKey := strings.TrimLeft(c.Param("key"), "/")
	if objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ключ объекта не указан"})
		return
	}

	err := services.DeleteFile(context.TODO(), []string{objectKey})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Файл удалён"})
}

func GetFileURLHandler(c *gin.Context) {
	objectKey := strings.TrimLeft(c.Param("key"), "/")
	if objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя файла не указано"})
		return
	}

	url := services.GetFileURL(context.TODO(), objectKey)
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func DownloadFileHandler(c *gin.Context) {
	objectKey := strings.TrimLeft(c.Param("key"), "/")
	if objectKey == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Имя файла не указано"})
		return
	}
	data, err := services.DownloadFile(context.TODO(), objectKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Data(http.StatusOK, "application/octet-stream", data)
}

func ListFilesHandler(c *gin.Context) {
	prefix := c.Query("prefix")
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
	c.JSON(http.StatusOK, gin.H{"buckets": res})
}
