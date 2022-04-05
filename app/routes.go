package app

import (
	"backend-onboarding/config"
	"backend-onboarding/delivery/product_delivery"
	"backend-onboarding/delivery/role_delivery"
	"backend-onboarding/delivery/user_delivery"
	"backend-onboarding/middleware"
	"backend-onboarding/repository/product_repository"
	"backend-onboarding/repository/role_repository"
	"backend-onboarding/repository/user_repository"
	"backend-onboarding/usecase/product_usecase"
	"backend-onboarding/usecase/role_usecase"
	"backend-onboarding/usecase/user_usecase"
	"backend-onboarding/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(mysqlConn *gorm.DB) *gin.Engine {
	config.InitConfig()

	userRepository := user_repository.GetUserRepository(mysqlConn)
	jwtAuth := user_usecase.GetJwtUsecase(userRepository)
	userUseCase := user_usecase.InsertUseCase(userRepository, jwtAuth)
	userDelivery := user_delivery.GetUserDelivery(userUseCase)

	productRepository := product_repository.GetProductRepository(mysqlConn)
	productUseCase := product_usecase.GetProductUseCase(productRepository)
	productDelivery := product_delivery.GetProductDelivery(productUseCase)

	roleRepository := role_repository.GetRoleRepository(mysqlConn)
	roleUseCase := role_usecase.GetRoleUseCase(roleRepository)
	roleDelivery := role_delivery.GetRoleDelivery(roleUseCase)
	defaultCors := utils.CORSMiddleware()

	router := gin.Default()
	router.Use(defaultCors)

	protectedRoutes := router.Group("/")
	protectedRoutes.Use(middleware.JWTAuth(jwtAuth))
	{
		protectedRoutes.GET("/users", userDelivery.UserList)
		protectedRoutes.GET("/user/:id", userDelivery.UserDetailById)
		protectedRoutes.GET("/products", productDelivery.GetAllProducts)
		protectedRoutes.GET("/product/:id", productDelivery.GetProductById)
		protectedRoutes.POST("/product", productDelivery.CreateNewProduct)
		protectedRoutes.GET("/roles", roleDelivery.RoleList)
	}

	adminRoutes := router.Group("/")
	adminRoutes.Use(middleware.JWTAuthAdmin(jwtAuth, userRepository))
	{
		adminRoutes.PUT("/user/:id", userDelivery.UpdateUserData)
		adminRoutes.DELETE("/user/:id", userDelivery.DeleteUserById)
		adminRoutes.PUT("/product/:id", productDelivery.UpdateProductById)
		adminRoutes.DELETE("/product/:id", productDelivery.DeleteProductById)

	}

	checkerRoutes := router.Group("/")
	checkerRoutes.Use(middleware.JWTAuthChecker(jwtAuth, userRepository))
	{
		checkerRoutes.PUT("/product/:id/checked", productDelivery.UpdateCheckedProduct)
	}

	signerRoutes := router.Group("/")
	signerRoutes.Use(middleware.JWTAuthSigner(jwtAuth, userRepository))
	{
		signerRoutes.PUT("/product/:id/signer", productDelivery.UpdatePublishedProduct)
	}

	router.POST("/user", userDelivery.InsertNewUser)
	router.POST("/login", userDelivery.UserLogin)

	router.Run(":8080")

	return router
}
