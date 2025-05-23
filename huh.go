package main

import (
	"fmt"
	"strconv"

	"github.com/AndrewMeleka/currency-converter/currency"
	"github.com/charmbracelet/huh"
)

func starterForm() *huh.Form {
	// Create a new form
	cOptions := currency.HuhOptions()
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Base Currency").
				Options(*cOptions...).
				Value(&from),
			huh.NewSelect[string]().
				Title("Target Currency").
				Options(*cOptions...).
				Value(&to),
			huh.NewInput().
				Title("Enter the amount?").
				Value(&amountStr).
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

	return form
}
