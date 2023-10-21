package route

import (
	"aime-api/config"
	"aime-api/log"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	serverAddr = ""
)

func Start(cfg *config.General) {
	if serverAddr != "" {
		return
	}

	serverAddr = cfg.Address
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(Cors())

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to AIME API",
			})
		})

		apiGroup.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		apiGroup.GET("/timestamp", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"timestamp": time.Now().Unix(),
			})
		})

		authGroup := apiGroup.Group("/auth")
		{
			authGroup.POST("/telegram", telegramOauth(cfg))
		}
	}

	l, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Errorln("External controller listen error: %s", err)
		return
	}
	serverAddr = l.Addr().String()
	log.Infoln("RESTful API listening at: %s", serverAddr)
	if err = http.Serve(l, r); err != nil {
		log.Errorln("External controller serve error: %s", err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range c.Request.Header {
			headerKeys = append(headerKeys, k)
		}

		var headerStr string
		headerStr = strings.Join(headerKeys, ", ")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			origin := c.Request.Header.Get("Origin")
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			c.Header("Access-Control-Max-Age", "172800")
			c.Header("Access-Control-Allow-Credentials", "false")
			c.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}
