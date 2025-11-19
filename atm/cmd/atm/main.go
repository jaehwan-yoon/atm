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
	//test card
	card := models.Card{
		Number:     "1234567890",
		HolderName: "JaeHwan Yoon",
	}
	accounts := []models.Account{
		{Number: "AC00001", Balance: 100000},
		{Number: "AC00002", Balance: 200000},
	}
	//0. add card and insert for mock server
	mockBank.AddCard(card.Number, "1234", accounts)
	mockHardware.InsertCard(&card)

	controller := atm.NewController(mockBank, mockHardware)
	err := controller.IsReadyAtmController(controller)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("===simulation start====")
	//1. intert
	fmt.Println("==1.Insert Card...")
	err = controller.InsertCard()
	if err != nil {
		fmt.Println("[Fail] Insert Card: ", err)
		return
	}
	//2. pin 입력
	fmt.Println("==2. Input Pin...")
	err = controller.EnterPin("1234")
	if err != nil {
		fmt.Println("[Fail] Enter Pin: ", err)
		return
	}

}
