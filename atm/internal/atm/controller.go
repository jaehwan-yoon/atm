package atm

import (
	"atm/internal/bank"
	"atm/pkg/models"
	"fmt"
)

type Controller struct {
	bankService bank.Service
	//card, account
	currentCard     models.Card
	selectedAccount models.Account
}

func NewController() *Controller {
	fmt.Println("NewController init")
	return &Controller{
		bankService: bank.NewMockService(),
	}
}

func (c *Controller) IsReadyAtmController(controller *Controller) error {
	fmt.Println("Atm Controller is ready")
	return nil
}
