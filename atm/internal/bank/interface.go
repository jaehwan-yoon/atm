package bank

import "atm/pkg/models"

type Service interface {
	VerifyPin(cardNumber string, pin string) (bool, error)
	GetAccounts(cardNumber string) ([]models.Account, error)
	GetBalance(accountNumber string) (int, error)
	Deposit(accountNumber string, amount int) error
	Withdraw(accountNumber string, amount int) error
}
