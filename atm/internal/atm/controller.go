package atm

import (
	"atm/pkg/models"
	"fmt"
)

type Controller struct {

	//card, account
	currentCard     models.Card
	selectedAccount models.Account
}

func NewController() *Controller {
	fmt.Println("NewController init")
	return &Controller{}
}

func (c *Controller) IsReadyAtmController(controller *Controller) error {
	fmt.Println("Atm Controller is ready")
	return nil
}
