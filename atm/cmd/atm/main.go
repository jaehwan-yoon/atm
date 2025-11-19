package main

import (
	"atm/internal/atm"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	controller := atm.NewController()
	err := controller.IsReadyAtmController(controller)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

}
