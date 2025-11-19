package atm

import (
	"atm/internal/bank"
	"atm/internal/hardware"
	"atm/pkg/models"
	"fmt"
)

type Controller struct {
	bankService     bank.Service
	hardwareService hardware.Service
	//card, account
	currentCard     models.Card
	selectedAccount models.Account
}

func NewController(bankService bank.Service, hardwareService hardware.Service) *Controller {
	fmt.Println("NewController init")
	return &Controller{
		bankService:     bank.NewMockService(),
		hardwareService: hardware.NewMockService(),
	}
}

func (c *Controller) IsReadyAtmController(controller *Controller) error {
	fmt.Println("Atm Controller is ready")
	return nil
}
