package endpoints

import (
	"net/http"

	"github.com/Rekka-Technologies/mabel/backend/internal/models"
	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	user_id, err := models.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := models.GetUserByID(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
