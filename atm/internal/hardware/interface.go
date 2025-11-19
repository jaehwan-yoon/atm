package hardware

import "atm/pkg/models"

type Service interface {
	ReadCard() (*models.Card, error)
}
