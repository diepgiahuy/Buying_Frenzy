package api

import (
	"github.com/diepgiahuy/Buying_Frenzy/docs"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/api/middleware"
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/diepgiahuy/Buying_Frenzy/util/config"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

type GinServerMode int

const (
	DebugMode GinServerMode = iota
	ReleaseMode
	TestMode
)

// GinServer : the struct gathering all the server details
type GinServer struct {
	config *config.ServerConfig
	router *gin.Engine
	store  *storage.PostgresStore
}

func initSwagger() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Buying Frenzy API"
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Description = "Backend service and a database for a food delivery platform"
	docs.SwaggerInfo.Host = ""
	docs.SwaggerInfo.BasePath = "/api/v1"
}

func NewServer(cfg *config.ServerConfig, store *storage.PostgresStore) *GinServer {
	s := GinServer{}
	s.store = store
	s.config = cfg
	mode := GinServerMode(cfg.Env)
	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}
	s.setupRouter()
	initSwagger()
	return &s
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept, Origin, Cache-Control, X-Requested-With, User-Agent, Accept-Language, Accept-Encoding")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// setupRouter
func (s *GinServer) setupRouter() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("api/v1")
	v1.GET("/restaurants", s.listRestaurantsOpen)
	v1.GET("/restaurants/top-list-with-price", s.listRestaurantsWithComparison)
	v1.GET("/restaurants/:name", s.listRestaurantsByName)
	v1.GET("/restaurants/dish/:name", s.listDishByName)
	v1.POST("/purchase", middleware.DBTransactionMiddleware(s.store.Db), s.createOrder)

	s.router = router
}

// Start the server
func (s GinServer) Start() {

	port := os.Getenv("PORT")
	if port == "" {
		port = s.config.Port // Default port if not specified
	}
	err := s.router.Run(":" + port)
	if err != nil {
		log.Fatal("Error during start server: ", err)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
