package infra

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

type GinServerMode int

const (
	DebugMode GinServerMode = iota
	ReleaseMode
	TestMode
)

// GinServer : the struct gathering all the server details
type GinServer struct {
	port   int
	Router *gin.Engine
}

// NewServer
func NewServer(port int, mode GinServerMode) GinServer {
	s := GinServer{}
	s.port = port

	s.Router = gin.New()

	switch mode {
	case DebugMode:
		gin.SetMode(gin.DebugMode)
	case TestMode:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.ReleaseMode)
	}

	s.Router.Use(gin.Recovery())
	s.Router.Use(CORSMiddleware())

	return s
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

// Start the server
func (s GinServer) Start() {
	s.Router.Run(":" + strconv.Itoa(int(s.port)))
}
