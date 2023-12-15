package handler

import (
	"github.com/gin-gonic/gin"
	"go-backer-api/helper"
	"go-backer-api/transaction"
	"go-backer-api/user"
	"net/http"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (t *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := t.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Campaign Transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	//response := helper.APIResponse("Campaign's transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)

}
