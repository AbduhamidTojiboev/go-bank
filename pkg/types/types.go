package types

type Money int

type Currency string

type PAN string

type Category string

type Status string

const (
	TJS       Currency = "TJS"
	RUB       Currency = "RUB"
	USD       Currency = "USD"
	TYPE_CARD string   = "card"
)

const (
	StatusOk         Status = "OK"
	StatusFail       Status = "Fail"
	StatusInProgress Status = "INPROGRESS"
)

type Card struct {
	ID       int
	PAN      PAN
	Currency Currency
	Color    string
	Balance  Money
	Name     string
	Active   bool
}

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

type PaymentSource struct {
	Type    string
	Number  PAN
	Balance Money
}
