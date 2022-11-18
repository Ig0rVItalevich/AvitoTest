package handler

import (
	"github.com/Ig0rVItalevich/avito-test/pkg/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Ig0rVItalevich/avito-test/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.GET("/:id", h.getUserBalance)
			users.POST("/refill", h.refillUserBalance)
		}

		api.POST("/transfer", h.transfer)

		purchase := api.Group("/purchase")
		{
			purchase.POST("/reserve", h.reservePurchase)
			purchase.POST("/accept", h.acceptPurchase)
			purchase.POST("/cancel", h.cancelPurchase)
		}

		reports := api.Group("/reports")
		{
			reports.GET("/user", h.getReportUser)
			reports.GET("/revenue", h.getReportRevenue)
		}
	}

	return router
}
