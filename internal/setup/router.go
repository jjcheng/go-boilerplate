package setup

import (
	"github.com/jjcheng/go-boilerplate/internal/cfg"
	"github.com/jjcheng/go-boilerplate/internal/middleware"
	"github.com/jjcheng/go-boilerplate/internal/service"
	"github.com/jjcheng/go-boilerplate/internal/types"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes and returns a Gin router with middleware and basic routes
func SetupRouter(logger *service.Logger) *gin.Engine {
	router := gin.New()
	router.Use(middleware.RequestID())
	router.Use(middleware.CORS())
	// authenticate middleware is added at router setup
	err := router.SetTrustedProxies(nil)
	if err != nil {
		panic(err.Error())
	}
	if cfg.Default().Site.Environment != types.EnvironmentDevelop {
		router.Use(middleware.Recovery(logger))
	}
	router.Use(middleware.Log(logger))
	pprof.Register(router)
	return router
}
