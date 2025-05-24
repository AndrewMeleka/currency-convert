package main

import (
	"fmt"
	"log"

	"github.com/AndrewMeleka/currency-converter/currency"
	"github.com/charmbracelet/huh"
	"github.com/inancgumus/screen"
)

var (
	amount    float64
	amountStr string
	from      string
	to        string
	again     bool
)

func main() {
	screen.Clear()
	screen.MoveTopLeft()

	starterForm := starterForm()

	if err := starterForm.Run(); err != nil {
		log.Fatal(err)
	}

	if from == to {
		fmt.Println("The currencies are the same, no conversion needed.")
		return
	}

	rate, err := currency.CurrenciesList.GetExchangeRate(from, to)
	if err != nil {
		fmt.Println("Error fetching exchange rate:", err)
	}

	fromC, _ := currency.CurrenciesList.GetCurrency(from)
	toC, _ := currency.CurrenciesList.GetCurrency(to)
	fmt.Printf("%s ➡️ %s\n", fromC.Code, toC.Code)

	fmt.Printf("The exchange rate : %.2f\n", rate)
	cAmount := amount * rate

	fmt.Printf("%s%.2f = %s%.2f\n", fromC.Symbol, amount, toC.Symbol, cAmount)

	err = huh.NewConfirm().
		Title("Would you like to try again?").
		Value(&again).
		Run()
	if err != nil {
		log.Fatal(err)
	}
	if again {
		main()
	} else {
		fmt.Println("Goodbye!")
	}
}
