package endpoints

import (
	"net/http"
	"time"

	"github.com/Rekka-Technologies/mabel/backend/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type TransactionInput struct {
	Name      string  `json:"name" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	Reference string  `json:"reference"`
	Date      string  `json:"date"`
}

func AddTransaction(c *gin.Context) {
	// Check if the user is authenticated
	user_id, err := models.ExtractTokenUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input
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

	// Set the date to today if it is not provided, this just allows the user
	// to override or backdate the transaction date.
	if input.Date == "" {
		t.Date = datatypes.Date(time.Now())
	} else {
		// Parse the date string as `YYYY-MM-DD` and set it as the transaction
		// date to backdate the transaction.
		backdate, err := time.Parse(time.DateOnly, input.Date)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		t.Date = datatypes.Date(backdate)
	}

	res, err := t.Create()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": res})
}

func GetTransactions(c *gin.Context) {
	// Check if the user is authenticated
	user_id, err := models.ExtractTokenUID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the transactions for the user
	t, err := models.GetTransactions(user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success", "data": t})
}
