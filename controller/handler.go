package controller

import (
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/lzj09/chat-server/controller/user"
	"github.com/lzj09/chat-server/server/ws"
	"github.com/lzj09/chat-server/utils"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
	"time"
)

func NewServerHandler() http.Handler {
	handler := gin.Default()
	handler.Use(Cors())

	// 开启ws监听
	go ws.WebsocketManager.Start()

	handler.GET("/health", func(context *gin.Context) {
		context.JSON(http.StatusOK, "health")
	})

	handler.POST("/v1/user/login", user.Login)

	handler.Use(Authenticator())

	handler.GET("/v1/ws", ws.Run)

	handler.GET("/v1/user/info", user.Info)

	return handler
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:        true,
		AllowOrigins:           nil,
		AllowOriginFunc:        nil,
		AllowMethods:           []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:           []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials:       true,
		ExposeHeaders:          nil,
		MaxAge:                 12 * time.Hour,
		AllowWildcard:          true,
		AllowBrowserExtensions: true,
		AllowWebSockets:        true,
		AllowFiles:             false,
	})
}

func Authenticator(pass ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 放行指定的路径
		for _, path := range pass {
			if path == c.Request.URL.Path {
				c.Next()
				return
			}
		}

		token, err := GetToken(c)
		if err != nil {
			klog.Errorf("get token error: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			c.Abort()
			return
		}

		t, err := jwt.ParseWithClaims(token, &user.TokenClaims{}, func() jwt.Keyfunc {
			return func(t *jwt.Token) (interface{}, error) {
				return []byte(utils.AccessSecret), nil
			}
		}())

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}

		if claims, ok := t.Claims.(*user.TokenClaims); ok && t.Valid {
			c.Set(utils.JWTPayloadKey, claims.ID)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			c.Abort()
			return
		}
	}
}

// GetToken 获取token
func GetToken(c *gin.Context) (string, error) {
	auth := c.GetHeader("Authorization")
	prefix := "Bearer "
	token := ""

	if auth != "" && strings.HasPrefix(auth, prefix) {
		token = auth[len(prefix):]
	} else {
		token = c.Request.FormValue("access_token")
		if token == "" {
			token = c.Request.FormValue("token")
		}
	}

	if token == "" {
		return "", errors.New("invalid access token")
	}

	return token, nil
}
