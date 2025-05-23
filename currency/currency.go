package currency

import (
	"fmt"

	"github.com/AndrewMeleka/currency-converter/util"
	"github.com/charmbracelet/huh"
)

type Currency string

var Currencies = map[Currency]string{
	"USD": "United States Dollar",
	"EUR": "Euro",
	"EGP": "Egyptian Pound",
	"GBP": "British Pound",
}

func CurrenciesList() *[]Currency {
	orderedSlice := util.MapToOrderedSlice[Currency](Currencies)
	return &orderedSlice
}

func HuhOptions() *[]huh.Option[string] {
	cOptions := []huh.Option[string]{}
	cl := *CurrenciesList()

	for _, ckey := range cl {
		c, ok := Currencies[ckey]
		if !ok {
			continue
		}
		label := fmt.Sprintf("%s (%s)", ckey, c)
		cOptions = append(cOptions, huh.NewOption(label, string(ckey)))
	}
	return &cOptions
}

func IsSupportedCurrency(c Currency) bool {
	if _, ok := Currencies[c]; ok {
		return true
	}
	return false
}
