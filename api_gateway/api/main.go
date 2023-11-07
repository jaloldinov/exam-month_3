package api

import (
	"api_gateway/config"
	"api_gateway/pkg/logger"
	"api_gateway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))
	// router.Use(MaxAllowed(100))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	apiV1 := router.Group("/v1")

	// category
	apiV1.POST("/category/create", handlerV1.CreateCategory)
	apiV1.GET("/category/list", handlerV1.GetAllCategory)
	apiV1.GET("/category/get/:category_id", handlerV1.GetCategory)
	apiV1.PUT("/category/update/:category_id", handlerV1.UpdateCategory)
	apiV1.DELETE("/category/delete/:category_id", handlerV1.DeleteCategory)

	// product
	apiV1.POST("/product/create", handlerV1.CreateProduct)
	apiV1.GET("/product/list", handlerV1.GetAllProduct)
	apiV1.GET("/product/get/:product_id", handlerV1.GetProduct)
	apiV1.PUT("/product/update/:product_id", handlerV1.UpdateProduct)
	apiV1.DELETE("/product/delete/:product_id", handlerV1.DeleteProduct)

	// order
	apiV1.POST("/order/create", handlerV1.CreateOrder)
	apiV1.GET("/order/list", handlerV1.GetAllOrder)
	apiV1.GET("/order/get/:order_id", handlerV1.GetOrder)
	apiV1.PUT("/order/update/:order_id", handlerV1.UpdateOrder)
	apiV1.DELETE("/order/delete/:order_id", handlerV1.DeleteOrder)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}
