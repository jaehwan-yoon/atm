package main

import (
	"atm/internal/atm"
	"atm/internal/bank"
	"atm/internal/hardware"
	"atm/pkg/models"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	mockBank := bank.NewMockService()
	mockHardware := hardware.NewMockService()
	controller := atm.NewController(mockBank, mockHardware)
	err := controller.IsReadyAtmController(controller)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	//test card
	card := models.Card{
		Number:     "1234567890",
		HolderName: "JaeHwan Yoon",
	}
	accounts := []models.Account{
		{Number: "AC00001", Balance: 100000},
		{Number: "AC00002", Balance: 200000},
	}
	mockBank.AddCard(card.Number, "1234", accounts)
	mockHardware.InsertCard(&card)
	//insert a card

}
