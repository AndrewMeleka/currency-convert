package currency

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

type Currency string

var CurrenciesList = map[Currency]string{
	"USD": "United States Dollar",
	"EUR": "Euro",
	"EGP": "Egyptian Pound",
	"GBP": "British Pound",
}

func HuhOptions() *[]huh.Option[string] {
	cOptions := []huh.Option[string]{}
	for k, v := range CurrenciesList {
		label := fmt.Sprintf("%s (%s)", k, v)
		cOptions = append(cOptions, huh.NewOption(label, string(k)))
	}
	return &cOptions
}

func IsSupportedCurrency(c Currency) bool {
	if _, ok := CurrenciesList[c]; ok {
		return true
	}
	return false
}
