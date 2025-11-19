package hardware

import (
	"atm/pkg/models"
	"fmt"
	"sync"

	"github.com/pkg/errors"
)

type MockService struct {
	mu sync.RWMutex

	CurrentCard *models.Card

	availableCash int
}

func NewMockService() *MockService {
	return &MockService{}
}

func (m *MockService) InsertCard(card *models.Card) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.CurrentCard = card
	fmt.Println("Card inserted: ", card.Number)
	return nil
}

func (m *MockService) ReadCard() (*models.Card, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if m.CurrentCard == nil {
		return nil, errors.New("[Hardware] no card inserted")
	}
	//get available cash

	return m.CurrentCard, nil
}

func (m *MockService) AcceptCash(amount int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if amount <= 0 {
		return errors.New("[Hardware] invalid amount")
	}
	m.availableCash += amount
	fmt.Println("[Hardware] Cash accepted: ", amount)
	return nil
}

func (m *MockService) DispenseCash(amount int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if amount <= 0 {
		return errors.New("[Hardware] invalid amount")
	}
	if m.availableCash < amount {
		return errors.New(fmt.Sprintf("[Hardware] insufficient cash: %d < %d", m.availableCash, amount))
	}
	m.availableCash -= amount
	fmt.Println("[Hardware] Cash dispensed: ", amount, ", remaining cash: ", m.availableCash)
	return nil
}

func (m *MockService) SetAvailableCash(amount int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if amount <= 0 {
		return errors.New("[Hardware] invalid amount")
	}
	m.availableCash = amount
	fmt.Println("[Hardware] Available cash set: ", amount)
	return nil
}
