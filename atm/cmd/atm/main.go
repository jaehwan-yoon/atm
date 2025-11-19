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
	fmt.Println("=card inserted")
	//2. pin pin
	fmt.Println("==2. Input Pin...")
	err = controller.EnterPin("1234")
	if err != nil {
		fmt.Println("[Fail] Enter Pin: ", err)
		return
	}
	fmt.Println("=pin verified")
	//3. accounts query
	fmt.Println("==3. Query Accounts...")
	availableAccounts, err := controller.GetAvailableAccounts()
	if err != nil {
		fmt.Println("[Fail] Query Accounts: ", err)
		return
	}
	fmt.Println("=accounts queried: ", availableAccounts)
	//4. accounts select
	fmt.Println("==4. Select Account...")
	err = controller.SelectAccount("AC00002")
	if err != nil {
		fmt.Println("[Fail] Select Account: ", err)
		return
	}
	fmt.Println("=account selected")

	//5. balace check
	fmt.Println("==5. Check Balance...")
	balance, err := controller.CheckBalance()
	if err != nil {
		fmt.Println("[Fail] Check Balance: ", err)
		return
	}
	fmt.Println("=balance checked: ", balance)

	//6. deposit money
	fmt.Println("==6. Deposit Money...")
	err = controller.Deposit(10000)
	if err != nil {
		fmt.Println("[Fail] Deposit: ", err)
		return
	}
	fmt.Println("=money deposited")

	//7. balance check
	fmt.Println("==7. Check Balance...")
	balance, err = controller.CheckBalance()
	if err != nil {
		fmt.Println("[Fail] Check Balance: ", err)
		return
	}
	fmt.Println("=balance checked: ", balance)

	//8. withdraw money
	fmt.Println("==8. Withdraw Money...")
	err = controller.Withdraw(30000)
	if err != nil {
		fmt.Println("[Fail] Withdraw: ", err)
		return
	}
	fmt.Println("=money withdrawn")

	//9. balance check
	fmt.Println("==9. Check Balance...")
	balance, err = controller.CheckBalance()
	if err != nil {
		fmt.Println("[Fail] Check Balance: ", err)
		return
	}
	fmt.Println("=balance checked: ", balance)

	//10. eject card
	fmt.Println("==10. Eject Card...")
	err = controller.EjectCard()
	if err != nil {
		fmt.Println("[Fail] Eject Card: ", err)
		return
	}
	fmt.Println("=card ejected")

	fmt.Println("===simulation end====")

}
