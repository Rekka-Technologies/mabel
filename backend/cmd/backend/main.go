package main

import (
	"net/http"

	auth "github.com/Rekka-Technologies/mabel/backend/internal/auth"
	"github.com/Rekka-Technologies/mabel/backend/internal/endpoints"
	"github.com/Rekka-Technologies/mabel/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", auth.Register)
	public.POST("/login", auth.Login)

	protected := r.Group("/api/admin")
	protected.Use(JwtAuthMiddleware())
	protected.GET("/user", endpoints.CurrentUser)
	protected.POST("/transactions", endpoints.AddTransaction)
	protected.GET("/transactions", endpoints.GetTransactions)

	r.Run(":8080")

}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := models.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
