package api

import (
	"github.com/diepgiahuy/Buying_Frenzy/pkg/storage"
	"github.com/diepgiahuy/Buying_Frenzy/util/config"
	"github.com/gin-gonic/gin"
	"log"
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
	store  *storage.Repo
}

func NewServer(cfg *config.ServerConfig, store *storage.Repo) *GinServer {
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
func (s *GinServer) setupRouter() {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(CORSMiddleware())
	router.Use(gin.Logger())
	v1 := router.Group("api/v1")
	v1.GET("/restaurants", s.listRestaurantsOpen)
	v1.GET("/restaurants/more-dishes", s.listRestaurantsWithMoreDishes)
	v1.GET("/restaurants/less-dishes", s.listRestaurantsWithLessDishes)
	v1.GET("/restaurants/:name", s.listRestaurantsByName)
	v1.GET("/restaurants/dish/:name", s.listRestaurantsByDishName)

	s.router = router
}

// Start the server
func (s GinServer) Start() {
	err := s.router.Run(":" + s.config.Port)
	if err != nil {
		log.Fatal("Error during start server: ", err)
	}
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
