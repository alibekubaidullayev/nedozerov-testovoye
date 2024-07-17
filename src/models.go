package main

import (
	"errors"
)

type BankAccount interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	GetBalance() float64
}

type AccountDetails struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

type Account struct {
	ID      int     `json:"id"`
	Balance float64 `json:"balance"`
	AccountDetails
}

func (a *Account) Deposit(amount float64) error {
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	if a.Balance < amount {
		return errors.New("Not enough balance")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	result := roundToTwoDecimalPlaces(a.Balance)
	return result
}
