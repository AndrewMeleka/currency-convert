package currency

import (
	"github.com/charmbracelet/huh"
)

var CurrenciesList = CurrencyList{
	Currencies: []Currency{
		{Code: "USD", Text: "United States Dollar", Symbol: "$"},
		{Code: "EUR", Text: "Euro", Symbol: "€"},
		{Code: "GBP", Text: "British Pound Sterling", Symbol: "£"},
		{Code: "EGP", Text: "Egyptian Pound", Symbol: "ج.م"},
	},
}

type CurrencyList struct {
	Currencies []Currency
}

func (cl *CurrencyList) GetCurrency(code string) (Currency, bool) {
	for _, c := range cl.Currencies {
		if c.Code == code {
			return c, true
		}
	}
	return Currency{}, false
}

func (cl *CurrencyList) HuhOptions() *[]huh.Option[string] {
	options := []huh.Option[string]{}
	for _, c := range cl.Currencies {
		options = append(options, huh.NewOption(c.String(), c.Code))
	}
	return &options
}
