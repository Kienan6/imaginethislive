package gin

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"itl/interceptor"
)

type Engine struct {
	Router *gin.Engine
}

type EngineParams struct {
	fx.In
	Middleware  []interceptor.Middleware                     `group:"interceptors"`
	Controllers []func(rg *gin.RouterGroup) *gin.RouterGroup `group:"controllers"`
}

func setupCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:9004"}
	router.Use(cors.New(config))
}

// NewEngine
// Setup gin instance with controllers/middleware
func NewEngine(params EngineParams) *Engine {
	router := gin.Default()
	setupCors(router)

	for _, i := range params.Middleware {
		router.Use(i.Run())
	}

	group := router.Group("/v1")

	for _, c := range params.Controllers {
		c(group)
	}

	return &Engine{
		Router: router,
	}
}
