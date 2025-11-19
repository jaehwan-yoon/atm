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
	currentCard       *models.Card
	selectedAccount   *models.Account
	availableAccounts []models.Account
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

func (c *Controller) EnterPin(pin string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.currentCard == nil {
		return errors.New("[Controller] No card inserted")
	}
	//verfy pin
	valid, err := c.bankService.VerifyPin(c.currentCard.Number, pin)
	if err != nil {
		return err
	}
	if !valid {
		return errors.New("[Controller] Invalid pin")
	}
	//get accounts
	accounts, err := c.bankService.GetAccounts(c.currentCard.Number)
	if err != nil {
		return err
	}
	if len(accounts) == 0 {
		return errors.New("[Controller] No accounts available")
	}
	c.availableAccounts = accounts
	fmt.Println("[Controller] Accounts: ", accounts)
	return nil
}

func (c *Controller) GetAvailableAccounts() ([]models.Account, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.availableAccounts == nil {
		return nil, errors.New("[Controller] No accounts available")
	}
	//return a depp copy
	results := make([]models.Account, len(c.availableAccounts))
	for i, account := range c.availableAccounts {
		results[i] = account
	}
	return results, nil
}

func (c *Controller) SelectAccount(accountNumber string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.selectedAccount != nil {
		return errors.New("[Controller] Account already selected")
	}
	//find account
	for _, account := range c.availableAccounts {
		if account.Number == accountNumber {
			c.selectedAccount = &account
			//hardware set available cash
			err := c.hardwareService.SetAvailableCash(account.Balance)
			if err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("[Controller] Account not found")
}

func (c *Controller) CheckBalance() (int, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.selectedAccount == nil {
		return 0, errors.New("[Controller] No account selected")
	}
	balance, err := c.bankService.GetBalance(c.selectedAccount.Number)
	if err != nil {
		return 0, err
	}
	c.selectedAccount.Balance = balance
	return balance, nil
}

func (c *Controller) Deposit(amount int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if amount <= 0 {
		return errors.New("[Controller] Invalid amount")
	}
	if c.selectedAccount == nil {
		return errors.New("[Controller] No account selected")
	}
	//accept cash from hardware
	err := c.hardwareService.AcceptCash(amount)
	if err != nil {
		return err
	}
	//deposit to bank
	err = c.bankService.Deposit(c.selectedAccount.Number, amount)
	if err != nil {
		return err
	}
	//update local account balance
	return nil
}

func (c *Controller) Withdraw(amount int) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if amount <= 0 {
		return errors.New("[Controller] Invalid amount")
	}
	if c.selectedAccount == nil {
		return errors.New("[Controller] No account selected")
	}

	//withdraw from bank first
	err := c.bankService.Withdraw(c.selectedAccount.Number, amount)
	if err != nil {
		return err
	}
	//dispense cash from hardware
	err = c.hardwareService.DispenseCash(amount)
	if err != nil {
		//todo: add rollback when dispense fails
		return err
	}

	//update local account balance
	balance, err := c.bankService.GetBalance(c.selectedAccount.Number)
	if err != nil {
		fmt.Println("[Controller] Get balance failed: ", err)
		return err
	}
	c.selectedAccount.Balance = balance
	fmt.Println("[Controller] Withdraw accountNumber: ", c.selectedAccount.Number, ", withdraw amount: ", amount, ", new balance: ", balance)
	return nil
}
