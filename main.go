package main

import (
	"fmt"
	"log"

	"github.com/AndrewMeleka/currency-converter/currency"
	"github.com/charmbracelet/huh"
	"github.com/inancgumus/screen"
)

var currenciesList = currency.List{
	Currencies: []currency.Currency{
		{Code: "USD", Text: "United States Dollar", Symbol: "$"},
		{Code: "EUR", Text: "Euro", Symbol: "€"},
		{Code: "GBP", Text: "British Pound Sterling", Symbol: "£"},
		{Code: "EGP", Text: "Egyptian Pound", Symbol: "E£"},
	},
}

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

	rate, err := currenciesList.GetExchangeRate(from, to)
	if err != nil {
		fmt.Println("Error fetching exchange rate:", err)
	}

	fromC, _ := currenciesList.GetCurrency(from)
	toC, _ := currenciesList.GetCurrency(to)
	fmt.Printf("%s ➡️ %s\n", fromC.Code, toC.Code)

	fmt.Printf("The exchange rate: %s%.2f\n", toC.Symbol, rate)
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
