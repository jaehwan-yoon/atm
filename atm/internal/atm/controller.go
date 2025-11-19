package atm

import (
	"atm/internal/bank"
	"atm/internal/hardware"
	"atm/pkg/models"
	"errors"
	"fmt"
	"sync"
)

type Controller struct {
	mu              sync.RWMutex
	bankService     bank.Service
	hardwareService hardware.Service
	//card, account
	currentCard     *models.Card
	selectedAccount *models.Account
}

func NewController(bankService bank.Service, hardwareService hardware.Service) *Controller {
	fmt.Println("NewController init")
	return &Controller{
		bankService:     bankService,
		hardwareService: hardwareService,
	}
}

func (c *Controller) IsReadyAtmController(controller *Controller) error {
	fmt.Println("Atm Controller is ready")
	return nil
}

func (c *Controller) InsertCard() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.currentCard != nil {
		return errors.New("[Controller] Card already inserted")
	}
	card, err := c.hardwareService.ReadCard()
	if err != nil {
		return err
	}
	c.currentCard = card
	fmt.Println("[Controller] Card inserted: ", card.Number)
	return nil
}
