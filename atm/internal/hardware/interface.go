package hardware

import "atm/pkg/models"

type Service interface {
	ReadCard() (*models.Card, error)
	AcceptCash(amount int) error
	DispenseCash(amount int) error
	SetAvailableCash(amount int) error
}
