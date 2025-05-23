package currency

import (
	"fmt"
	"strconv"

	"github.com/AndrewMeleka/currency-converter/scrapper"
)

func GetExchangeRate(from Currency, to Currency) (float64, error) {
	if !IsSupportedCurrency(from) || !IsSupportedCurrency(to) {
		return 0, fmt.Errorf("unsupported currency: %s or %s", from, to)
	}
	url := fmt.Sprintf("https://wise.com/gb/currency-converter/%s-to-%s-rate", from, to)

	s := scrapper.Scrapper{
		Url:       url,
		QueryFile: "wise-query-currency.txt",
	}
	rate, err := s.Scrap()
	if err != nil {
		return 0, fmt.Errorf("failed to get exchange rate: %v", err)
	}
	rateFloat, err := strconv.ParseFloat(rate, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse exchange rate: %v", err)
	}
	return rateFloat, nil
}
