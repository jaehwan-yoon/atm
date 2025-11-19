package bank

import "atm/pkg/models"

type Service interface {
	VerifyPin(cardNumber string, pin string) (bool, error)
	GetAccounts(cardNumber string) ([]models.Account, error)
}
