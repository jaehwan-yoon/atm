package bank

import (
	"atm/pkg/models"
	"fmt"
	"sync"
)

type MockService struct {
	mu sync.RWMutex

	//card pins: card number -> pin
	pins map[string]string

	//card accounts: card number -> []accounts
	cardAccounts map[string][]models.Account

	//account balances: account number -> balance
	accountBalances map[string]int

	//accounts: account number -> account
	accounts map[string]models.Account
}

func NewMockService() *MockService {
	return &MockService{
		pins:            make(map[string]string),
		cardAccounts:    make(map[string][]models.Account),
		accountBalances: make(map[string]int),
		accounts:        make(map[string]models.Account),
	}
}

func (m *MockService) AddCard(cardNumber string, pin string, accounts []models.Account) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.pins[cardNumber] = pin
	m.cardAccounts[cardNumber] = accounts
	for _, account := range accounts {
		m.accountBalances[account.Number] = account.Balance
		m.accounts[account.Number] = account
	}
	fmt.Println("Card added: ", cardNumber)
}
