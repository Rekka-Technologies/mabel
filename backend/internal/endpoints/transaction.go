package endpoints

import (
	"net/http"

	"github.com/Rekka-Technologies/mabel/backend/internal/models"
	"github.com/gin-gonic/gin"
)

type TransactionInput struct {
	Name      string  `json:"name" binding:"required"`
	Reference string  `json:"reference"`
	Amount    float64 `json:"amount" binding:"required"`
}

func AddTransaction(c *gin.Context) {
	user_id, err := models.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var input TransactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t := models.Transaction{
		UserId:    user_id,
		Name:      input.Name,
		Reference: input.Reference,
		Amount:    input.Amount,
	}

	res, err := t.SaveTransaction()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": res})
}

func GetTransactions(c *gin.Context) {
	user_id, err := models.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	t, err := models.GetTransactions(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": t})
}
