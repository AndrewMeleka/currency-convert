package currency

import (
	"fmt"
)

type Currency struct {
	Code   string
	Text   string
	Symbol string
}

func (c Currency) String() string {
	return fmt.Sprintf("%s (%s)", c.Code, c.Text)
}
