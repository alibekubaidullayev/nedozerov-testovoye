package main

import (
	"errors"
	"sync"
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
	Mutex sync.Mutex
}

func (a *Account) Deposit(amount float64) error {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount float64) error {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	if a.Balance < amount {
		return errors.New("Not enough balance")
	}
	a.Balance -= amount
	return nil
}

func (a *Account) GetBalance() float64 {
	a.Mutex.Lock()
	defer a.Mutex.Unlock()
	result := roundToTwoDecimalPlaces(a.Balance)
	return result
}
