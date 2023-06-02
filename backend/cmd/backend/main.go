package main

import (
	"net/http"

	auth "github.com/Rekka-Technologies/mabel/backend/internal/auth"
	"github.com/Rekka-Technologies/mabel/backend/internal/endpoints"
	"github.com/Rekka-Technologies/mabel/backend/internal/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	models.ConnectDataBase()

	// Setup Router
	r := gin.Default()

	// Setup the CORS middleware
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	public := r.Group("/auth")
	public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)

	protected := r.Group("/api")
	protected.Use(JwtAuthMiddleware())
	protected.POST("/transactions", endpoints.AddTransaction)
	protected.GET("/transactions", endpoints.GetTransactions)

	r.Run(":8080")
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := models.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
