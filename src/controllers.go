package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var Mutex sync.Mutex

func createAcc(c *gin.Context) {
	var details AccountDetails

	if err := c.ShouldBindJSON(&details); err != nil {
		slog.Error("Error binding JSON", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error binding JSON"})
		return
	}

	newAcc := Account{
		AccountDetails: details,
		Balance:        0.0,
		ID:             len(db) + 1,
	}
	db = append(db, newAcc)

	slog.Info("Created Account: ", "account", newAcc)
	c.JSON(http.StatusOK, gin.H{"created-account": newAcc})
}

func deposit(c *gin.Context) {
	amount, index, err := readAmountAndId(c)
	if err != nil {
		slog.Error("Error reading id and amount", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Acc BankAccount = &db[index]

	errorChan := make(chan error)
	go func() {
		Mutex.Lock()
		defer Mutex.Unlock()
		errorChan <- Acc.Deposit(amount)
	}()

	if err := <-errorChan; err != nil {
		slog.Error("Deposit Error", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	successMsg := fmt.Sprintf("Successfully added %.2f to user with id %d", amount, index+1)
	slog.Info(fmt.Sprintf("ID: %d; Deposit: %.2f", index+1, amount))
	c.JSON(http.StatusOK, gin.H{"success": successMsg})
}

func withdraw(c *gin.Context) {
	amount, index, err := readAmountAndId(c)
	if err != nil {
		slog.Error("Error reading id and amount", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var Acc BankAccount = &db[index]

	errorChan := make(chan error)
	go func() {
		Mutex.Lock()
		defer Mutex.Unlock()
		errorChan <- Acc.Withdraw(amount)
	}()

	if err := <-errorChan; err != nil {
		slog.Error("Withdraw Error", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	successMsg := fmt.Sprintf("Successfully withdrawn %.2f from user with id %d", amount, index+1)
	slog.Info(fmt.Sprintf("ID: %d; Withdraw: %.2f", index+1, amount))
	c.JSON(http.StatusOK, gin.H{"success": successMsg})
}

func balance(c *gin.Context) {
	index, err := getIndex(c)
	if err != nil {
		slog.Error("Error index geting", "error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var Acc BankAccount = &db[index]

	floatChan := make(chan float64)
	go func() {
		Mutex.Lock()
		defer Mutex.Unlock()
		floatChan <- Acc.GetBalance()
	}()

	balance := <-floatChan
	slog.Info(fmt.Sprintf("ID: %d; Balance: %.2f", index+1, balance))
	c.JSON(http.StatusOK, gin.H{"balance": balance})
}
