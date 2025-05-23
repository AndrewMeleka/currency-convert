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

	rate, err := currency.GetExchangeRate(currency.Currency(from), currency.Currency(to))
	if err != nil {
		fmt.Println("Error fetching exchange rate:", err)
	}

	fmt.Printf("The exchange rate from %s to %s is: %.2f\n", from, to, rate)
	cAmount := amount * rate
	fmt.Printf("The converted amount (%.2f): %.2f\n", amount, cAmount)

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
