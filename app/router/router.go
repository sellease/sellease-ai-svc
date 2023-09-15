package router

import (
	"sellease-ai/internal/handler"
	"sellease-ai/internal/middleware"
	"sellease-ai/internal/repository"
	"sellease-ai/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RouterContext struct {
	// DB      *gorm.DB
	Repo    *repository.Repository
	Usecase *usecase.Usecase
	Handler *handler.Handler
}

func PrepareRouter(rc *RouterContext) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	router.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/sellease-ai/health"),
		gin.Recovery(),
	)

	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.Trace())

	app := router.Group("sellease-ai")
	rc.HealthRoutes(app)

	// user := app.Group("user")
	// rc.UserRoutes(user)

	prod := app.Group("prod")
	rc.ProductRoutes(prod)

	return router
}
