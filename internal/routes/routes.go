package routes

import (
	"STDE_proj/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	s3Group := router.Group("/s3")
	{
		// POST — загрузка файла
		s3Group.POST("/upload", controllers.PostFileHandler)

		// DELETE — удаление файла
		s3Group.DELETE("/files/:filename", controllers.DeleteFileHandler)

		// Список бакетов
		s3Group.GET("/buckets", controllers.ListBucketsHandler)

		// Список файлов
		s3Group.GET("/files", controllers.ListFilesHandler)

		// GET — получение файла/подписанного URL
		s3Group.GET("/files/url/:bucket/:filename", controllers.GetFileURLHandler)

		// Скачивание файла
		s3Group.GET("/download/:filename", controllers.DownloadFileHandler)

		// Пока не сделал
		//s3Group.PUT("/files/:filename", s3Controller.PutFileHandler)
	}
	//User groups
	router.GET("/user_groups", controllers.GetUserGroupsHandler)
	router.POST("/user_groups", controllers.PostUserGroupHandler)
	router.GET("/user_groups/:id", controllers.GetUserGroupByIdHandler)
	router.DELETE("/user_groups/:id", controllers.DeleteUserGroupHandler)
	router.PUT("/user_groups/:id", controllers.PutUserGroupHandler)
	//Permission
	router.GET("/permission", controllers.GetPermissionHandler)
	router.POST("/permission", controllers.PostPermissionHandler)
	router.GET("/permission/:id", controllers.GetPermissionByIdHandler)
	router.DELETE("/permission/:id", controllers.DeletePermissionHandler)
	router.PUT("/permission/:id", controllers.PutPermissionHandler)
	//Position
	router.GET("/position", controllers.GetPositionsHandler)
	router.POST("/position", controllers.PostPositionHandler)
	router.GET("/position/:id", controllers.GetPositionByIdHandler)
	router.DELETE("/position/:id", controllers.DeletePositionHandler)
	router.PUT("/position/:id", controllers.PutPositionHandler)

	router.POST("/auth_group_permissions", controllers.PostAuthGroupPermissionsHandler)
}
