package main

import (
	"errors"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

func roundToTwoDecimalPlaces(num float64) float64 {
	return math.Round(num*100) / 100
}

type bodyStruct struct {
	Amount float64 `json:"amount"`
}

func readAmountAndId(c *gin.Context) (float64, int, error) {
	body := bodyStruct{}
	if err := c.ShouldBindJSON(&body); err != nil {
		return 0, 0, errors.New("Error reading JSON")
	}

	index, err := getIndex(c)
	if err != nil {
		return 0, 0, err
	}

	return body.Amount, index, nil
}

func getIndex(c *gin.Context) (int, error) {
	id := c.Param("id")
	index, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("Invalid ID")
	}

	index--
	if err := validateID(index); err != nil {
		return 0, err
	}

	return index, nil
}

func validateID(index int) error {
	if index < 0 || index >= len(db) {
		return errors.New("No user with such ID")
	}
	return nil
}
