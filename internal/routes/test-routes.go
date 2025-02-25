package routes

import (
	"STDE_proj/internal/controllers"
	"github.com/gin-gonic/gin"
)

func TestRoutes(router *gin.Engine) {
	test := router.Group("api/test-routes")
	{

		s3Group := test.Group("/s3")
		{
			// POST — загрузка файла с поддержкой пути *key (можно писать /папка/../папка/файл.расширение)
			s3Group.POST("/upload/*key", controllers.PostFileHandler)

			// DELETE с поддержкой пути *key (можно писать /папка/../папка/файл.расширение)
			s3Group.DELETE("/files/*key", controllers.DeleteFileHandler)

			// Список бакетов
			s3Group.GET("/buckets", controllers.ListBucketsHandler)

			// Список файлов
			s3Group.GET("/files", controllers.ListFilesHandler)

			// GET — получение подписанного URL с поддержкой пути *key (можно писать /папка/../папка/файл.расширение)
			s3Group.GET("/files/url/*key", controllers.GetFileURLHandler)
			
			// Скачивание файла с поддержкой пути
			s3Group.GET("/download/*key", controllers.DownloadFileHandler)

		}

		// Создание вложенной группы маршрутов для "reviews" внутри группы "test"
		reviews := test.Group("/reviews")
		{
			// Обработчики маршрутов для "reviews"
			reviews.GET("/", controllers.GetReviewsAllHandler)
			reviews.GET("/:id", controllers.GetReviewsByIdHandler)
			reviews.POST("/", controllers.PostReviewsHandler)
		}
		authUserGroups := test.Group("/auth-user-groups")
		{
			// Auth user groups
			authUserGroups.GET("/", controllers.GetAuthUserGroupsAllHandler)
			authUserGroups.GET("/:id", controllers.GetAuthUserGroupsByIdHandler)
			authUserGroups.POST("/", controllers.PostAuthUserGroupsHandler)
			authUserGroups.PUT("/:id", controllers.PutAuthUserGroupsHandler)
			authUserGroups.DELETE("/:id", controllers.DeleteAuthUserGroupsHandler)
		}
		UserGroups := test.Group("/user-groups")
		{
			UserGroups.GET("/", controllers.GetUserGroupsHandler)
			UserGroups.POST("/", controllers.PostUserGroupHandler)
			UserGroups.GET("/:id", controllers.GetUserGroupByIdHandler)
			UserGroups.DELETE("/:id", controllers.DeleteUserGroupHandler)
			UserGroups.PUT("/:id", controllers.PutUserGroupHandler)
		}
		Permissions := test.Group("/permissions")
		{
			Permissions.GET("/", controllers.GetPermissionHandler)
			Permissions.POST("/", controllers.PostPermissionHandler)
			Permissions.GET("/:id", controllers.GetPermissionByIdHandler)
			Permissions.DELETE("/:id", controllers.DeletePermissionHandler)
			Permissions.PUT("/:id", controllers.PutPermissionHandler)
		}
		Position := test.Group("/position")
		{
			Position.GET("/", controllers.GetPositionsHandler)
			Position.POST("/", controllers.PostPositionHandler)
			Position.GET("/:id", controllers.GetPositionByIdHandler)
			Position.DELETE("/:id", controllers.DeletePositionHandler)
			Position.PUT("/:id", controllers.PutPositionHandler)
		}
		AuthGroupsPermissions := test.Group("/auth-groups-permissions")
		{
			AuthGroupsPermissions.GET("/", controllers.GetAuthGroupPermissionsHandler)
			AuthGroupsPermissions.GET("/:id", controllers.GetAuthGroupPermissionsIdHandler)
			AuthGroupsPermissions.POST("/", controllers.PostAuthGroupPermissionsHandler)
			AuthGroupsPermissions.PUT("/:id", controllers.PutAuthGroupPermissionsHandler)
			AuthGroupsPermissions.DELETE("/:id", controllers.DeleteAuthGroupPermissionsHandler)
		}
		UserProfile := test.Group("/user-profile")
		{
			UserProfile.GET("/", controllers.GetUserProfileHandler)
			UserProfile.GET("/:id", controllers.GetUserProfileByIdHandler)
			UserProfile.POST("/", controllers.PostUserProfileHandler)
			UserProfile.PUT("/:id", controllers.PutUserProfileHandler)
			UserProfile.DELETE("/:id", controllers.DeleteUserProfileHandler)
		}
		Basket := test.Group("/basket")
		{
			Basket.GET("/", controllers.GetBasketsHandler)
			Basket.GET("/:id", controllers.GetBasketByIdUserHandler)
			Basket.POST("/", controllers.PostBasketHandler)
			//Basket.PUT("/:id", controllers.PutBasketHandler)
			Basket.DELETE("/*id", controllers.DeleteBasketHandler) //Добавить "user/" для удаления всех позиций у пользователя
		}
		Menu := test.Group("/menu")
		{
			Menu.GET("/", controllers.GetMenuHandler)
			Menu.GET("/:id", controllers.GetMenuByIdHandler)
			Menu.POST("/", controllers.PostMenuHandler)
			Menu.PUT("/:id", controllers.PutMenuHandler)
			Menu.DELETE("/:id", controllers.DeleteMenuHandler)
		}
		MenuPosition := test.Group("/menu-position")
		{
			MenuPosition.GET("/", controllers.GetMenuPositionsHandler)
			MenuPosition.GET("/:id", controllers.GetMenuPositionByIdHandler)
			MenuPosition.POST("/", controllers.PostMenuPositionHandler)
			MenuPosition.PUT("/:id", controllers.PutMenuPositionHandler)
			MenuPosition.DELETE("/:id", controllers.DeleteMenuPositionHandler)
		}
		OrderPosition := test.Group("/order-position")
		{
			OrderPosition.GET("/", controllers.GetOrderPositionsHandler)
			OrderPosition.GET("/:id", controllers.GetOrderPositionByIdHandler)
			OrderPosition.POST("/", controllers.PostOrderPositionHandler)
			OrderPosition.PUT("/:id", controllers.PutOrderPositionHandler)
			OrderPosition.DELETE("/:id", controllers.DeleteOrderPositionHandler)
		}
		DocumentAuthUser := test.Group("/document-auth-user")
		{
			DocumentAuthUser.GET("/", controllers.GetDocumentAuthUsersHandler)
			DocumentAuthUser.GET("/:id", controllers.GetDocumentAuthUserByIdHandler)
			DocumentAuthUser.POST("/", controllers.PostDocumentAuthUserHandler)
			DocumentAuthUser.PUT("/:id", controllers.PutDocumentAuthUserHandler)
			DocumentAuthUser.DELETE("/:id", controllers.DeleteDocumentAuthUserHandler)
		}
		UserDocument := test.Group("/user-document")
		{
			UserDocument.GET("/", controllers.GetUserDocumentsHandler)
			UserDocument.GET("/:id", controllers.GetUserDocumentByIdHandler)
			UserDocument.POST("/", controllers.PostUserDocumentHandler)
			UserDocument.PUT("/:id", controllers.PutUserDocumentHandler)
			UserDocument.DELETE("/:id", controllers.DeleteUserDocumentHandler)
		}
	}
}
