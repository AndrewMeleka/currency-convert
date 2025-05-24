package currency

import (
	"github.com/charmbracelet/huh"
)

type List struct {
	Currencies []Currency
}

func (cl *List) GetCurrency(code string) (Currency, bool) {
	for _, c := range cl.Currencies {
		if c.Code == code {
			return c, true
		}
	}
	return Currency{}, false
}

func (cl *List) HuhOptions() *[]huh.Option[string] {
	options := []huh.Option[string]{}
	for _, c := range cl.Currencies {
		options = append(options, huh.NewOption(c.String(), c.Code))
	}
	return &options
}
