package main

import (
	"fmt"
	"log"
	"strconv"

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

	cOptions := currency.HuhOptions()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Base Currency").
				Options(
					*cOptions...,
				).
				Value(&from),
			huh.NewSelect[string]().
				Title("Target Currency").
				Options(*cOptions...).
				Value(&to),
			huh.NewInput().
				Title("Enter the amount?").
				Value(&amountStr).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(input string) error {
					// Try to parse the input into a float
					val, err := strconv.ParseFloat(input, 64)
					if err != nil {
						return fmt.Errorf("please enter a valid number")
					}
					if val <= 0 {
						return fmt.Errorf("amount must be greater than 0")
					}
					amount = val
					return nil
				}),
		),
	)

	if err := form.Run(); err != nil {
		log.Fatal(err)
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
