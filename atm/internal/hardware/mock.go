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
		return nil, errors.New("no card inserted")
	}
	return m.CurrentCard, nil
}
