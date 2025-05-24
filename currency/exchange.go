package currency

import (
	"fmt"
	"os"
	"strconv"

	"github.com/AndrewMeleka/currency-converter/scrapper"
)

func (cl *List) GetExchangeRate(from string, to string) (float64, error) {
	fromC, fromOk := cl.GetCurrency(from)
	toC, toOk := cl.GetCurrency(to)
	if !fromOk || !toOk {
		return 0, fmt.Errorf("unsupported currency: %s or %s", from, to)
	}
	url := fmt.Sprintf("https://wise.com/gb/currency-converter/%s-to-%s-rate", fromC.Code, toC.Code)
	selector := "#calculator > div.tapestry-wrapper span[dir='ltr'] span"
	if os.Getenv("selector") != "" {
		selector = os.Getenv("selector")
	}
	s := scrapper.Scrapper{
		Url:      url,
		Selector: selector,
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
