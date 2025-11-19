package bank

import (
	"atm/pkg/models"
	"errors"
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

func (m *MockService) VerifyPin(cardNumber string, pin string) (bool, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if _, ok := m.pins[cardNumber]; !ok {
		return false, errors.New("[Bank] Card not found")
	}
	if m.pins[cardNumber] != pin {
		return false, errors.New("[Bank] Invalid pin")
	}
	return true, nil
}

func (m *MockService) GetAccounts(cardNumber string) ([]models.Account, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	accounts, exists := m.cardAccounts[cardNumber]
	if !exists {
		return nil, errors.New("[Bank] Card not found")
	}
	results := make([]models.Account, len(accounts))
	for i, account := range accounts {
		results[i] = models.Account{
			Number:  account.Number,
			Balance: m.accountBalances[account.Number],
		}
	}
	return results, nil
}
