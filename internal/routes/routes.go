package routes

import (
    "STDE_proj/internal/controllers"
    "STDE_proj/internal/controllers/Auth"
    "STDE_proj/internal/controllers/RegisterController"
    "STDE_proj/internal/middleware"
    "github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

    //JWTSecret := os.Getenv("JWT_SECRET")

    public := router.Group("/api/public")
    {
        user := public.Group("/user")
        {
            user.GET("/logout", Auth.LogoutHandler)
            // Маршрут не реализован как регистрация. Реализован как обычный Create
            //user.POST("/add", controllers.PostAuthUserHandler)
            register := user.Group("/register")
            {
                register.POST("/", RegisterController.RegisterHandler)
            }

            user.POST("/verify/:type", controllers.VerifyControllerHandler)
        }

        public.POST("/auth", controllers.AuthenticationHandler)

        token := public.Group("/token")
        {
            token.POST("/", controllers.GenerateAccessRefreshToken)
            token.GET("/check", controllers.TokenCheckController)
            token.POST("/refresh", controllers.RefreshToken)
        }

        // Добавленные S3 маршруты в публичную зону
        s3Public := public.Group("/s3")
        {
            // GET — получение подписанного URL (публичный доступ)
            s3Public.GET("/files/url/*key", controllers.GetFileURLHandler)
            
            // Скачивание файла (публичный доступ)
            s3Public.GET("/download/*key", controllers.DownloadFileHandler)
            
            // Список бакетов (публичный доступ)
            s3Public.GET("/buckets", controllers.ListBucketsHandler)
            
            // Список файлов (публичный доступ)
            s3Public.GET("/files", controllers.ListFilesHandler)
        }
    }

    protected := router.Group("/api/private")
    protected.Use(middleware.AuthMiddleware())
    {
        user := protected.Group("/user")
        {
            user.DELETE("/delete/:id", controllers.DeleteAuthUserHandler)

            Basket := user.Group("/basket")
            {
                Basket.GET("/", controllers.GetBasketsHandler)
                Basket.GET("/:id", controllers.GetBasketByIdUserHandler)
                Basket.POST("/", controllers.PostBasketHandler)
                //Basket.PUT("/:id", controllers.PutBasketHandler)
                Basket.DELETE("/*id", controllers.DeleteBasketHandler) //Добавить "user/" для удаления всех позиций у пользователя
            }
        }

        reviews := protected.Group("/reviews")
        {
            reviews.GET("/", controllers.GetReviewsAllHandler)
            reviews.GET("/:id", controllers.GetReviewsByIdHandler)
            reviews.POST("/", controllers.PostReviewsHandler)
        }

        AuthUserGroups := protected.Group("/auth-user-groups")
        {
            AuthUserGroups.GET("/", controllers.GetAuthUserGroupsAllHandler)
            AuthUserGroups.GET("/:id", controllers.GetAuthUserGroupsByIdHandler)
            AuthUserGroups.POST("/", controllers.PostAuthUserGroupsHandler)
            AuthUserGroups.PUT("/:id", controllers.PutAuthUserGroupsHandler)
            AuthUserGroups.DELETE("/:id", controllers.DeleteAuthUserGroupsHandler)
        }

        UserGroups := protected.Group("/user-groups")
        {
            UserGroups.GET("/", controllers.GetUserGroupsHandler)
            UserGroups.POST("/", controllers.PostUserGroupHandler)
            UserGroups.GET("/:id", controllers.GetUserGroupByIdHandler)
            UserGroups.DELETE("/:id", controllers.DeleteUserGroupHandler)
            UserGroups.PUT("/:id", controllers.PutUserGroupHandler)
        }

        Permissions := protected.Group("/permissions")
        {
            Permissions.GET("/", controllers.GetPermissionHandler)
            Permissions.POST("/", controllers.PostPermissionHandler)
            Permissions.GET("/:id", controllers.GetPermissionByIdHandler)
            Permissions.DELETE("/:id", controllers.DeletePermissionHandler)
            Permissions.PUT("/:id", controllers.PutPermissionHandler)
        }

        Position := protected.Group("/position")
        {
            Position.GET("/", controllers.GetPositionsHandler)
            Position.POST("/", controllers.PostPositionHandler)
            Position.GET("/:id", controllers.GetPositionByIdHandler)
            Position.DELETE("/:id", controllers.DeletePositionHandler)
            Position.PUT("/:id", controllers.PutPositionHandler)
        }

        AuthGroupsPermissions := protected.Group("/auth-groups-permissions")
        {
            AuthGroupsPermissions.GET("/", controllers.GetAuthGroupPermissionsHandler)
            AuthGroupsPermissions.GET("/:id", controllers.GetAuthGroupPermissionsIdHandler)
            AuthGroupsPermissions.POST("/", controllers.PostAuthGroupPermissionsHandler)
            AuthGroupsPermissions.PUT("/:id", controllers.PutAuthGroupPermissionsHandler)
            AuthGroupsPermissions.DELETE("/:id", controllers.DeleteAuthGroupPermissionsHandler)
        }

        UserProfile := protected.Group("/user-profile")
        {
            UserProfile.GET("/", controllers.GetUserProfileHandler)
            UserProfile.GET("/:id", controllers.GetUserProfileByIdHandler)
            UserProfile.POST("/", controllers.PostUserProfileHandler)
            UserProfile.PUT("/:id", controllers.PutUserProfileHandler)
            UserProfile.DELETE("/:id", controllers.DeleteUserProfileHandler)
        }

        // Добавленные S3 маршруты в защищенную зону
        s3Protected := protected.Group("/s3")
        {
            // POST — загрузка файла (требуется аутентификация)
            s3Protected.POST("/upload/*key", controllers.PostFileHandler)
            
            // DELETE файла (требуется аутентификация)
            s3Protected.DELETE("/files/*key", controllers.DeleteFileHandler)
        }

        Menu := protected.Group("/menu")
        {
            Menu.GET("/", controllers.GetMenuHandler)
            Menu.GET("/:id", controllers.GetMenuByIdHandler)
            Menu.POST("/", controllers.PostMenuHandler)
            Menu.PUT("/:id", controllers.PutMenuHandler)
            Menu.DELETE("/:id", controllers.DeleteMenuHandler)
        }

        MenuPosition := protected.Group("/menu-position")
        {
            MenuPosition.GET("/", controllers.GetMenuPositionsHandler)
            MenuPosition.GET("/:id", controllers.GetMenuPositionByIdHandler)
            MenuPosition.POST("/", controllers.PostMenuPositionHandler)
            MenuPosition.PUT("/:id", controllers.PutMenuPositionHandler)
            MenuPosition.DELETE("/:id", controllers.DeleteMenuPositionHandler)
        }

        OrderPosition := protected.Group("/order-position")
        {
            OrderPosition.GET("/", controllers.GetOrderPositionsHandler)
            OrderPosition.GET("/:id", controllers.GetOrderPositionByIdHandler)
            OrderPosition.POST("/", controllers.PostOrderPositionHandler)
            OrderPosition.PUT("/:id", controllers.PutOrderPositionHandler)
            OrderPosition.DELETE("/:id", controllers.DeleteOrderPositionHandler)
        }

        DocumentAuthUser := protected.Group("/document-auth-user")
        {
            DocumentAuthUser.GET("/", controllers.GetDocumentAuthUsersHandler)
            DocumentAuthUser.GET("/:id", controllers.GetDocumentAuthUserByIdHandler)
            DocumentAuthUser.POST("/", controllers.PostDocumentAuthUserHandler)
            DocumentAuthUser.PUT("/:id", controllers.PutDocumentAuthUserHandler)
            DocumentAuthUser.DELETE("/:id", controllers.DeleteDocumentAuthUserHandler)
        }

        UserDocument := protected.Group("/user-document")
        {
            UserDocument.GET("/", controllers.GetUserDocumentsHandler)
            UserDocument.GET("/:id", controllers.GetUserDocumentByIdHandler)
            UserDocument.POST("/", controllers.PostUserDocumentHandler)
            UserDocument.PUT("/:id", controllers.PutUserDocumentHandler)
            UserDocument.DELETE("/:id", controllers.DeleteUserDocumentHandler)
        }
    }
}