package types

type Money int

type Currency string

type PAN string

const (
	TJS       Currency = "TJS"
	RUB       Currency = "RUB"
	USD       Currency = "USD"
	TYPE_CARD string   = "card"
)

type Card struct {
	ID         int
	PAN        PAN
	Currency   Currency
	Color      string
	Balance    Money
	Name       string
	Active     bool
	MinBalance Money
}

type Payment struct {
	ID     int
	Amount Money
}

type PaymentSource struct {
	Type    string
	Number  PAN
	Balance Money
}
