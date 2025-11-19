package hardware

import (
	"atm/pkg/models"
	"fmt"
	"sync"
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
