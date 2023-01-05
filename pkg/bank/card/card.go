package card

import (
	"github.com/AbduhamidTojiboev/go-bank/pkg/bank/types"
)

const withdrawLimit = 20_000_00
const depositLimit = 50_000_00
const bonusLimit = 5_000_00
const PercentBonus int = 3

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}

	return card
}

func Withdraw(card *types.Card, amount types.Money) {
	if validateWithdraw(*card, amount) {
		card.Balance -= amount
	}
}

func Deposit(card *types.Card, amount types.Money) {
	if validateDeposit(*card, amount) {
		card.Balance += amount
	}
}

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if !validateActiveAndBalance(*card) {
		return
	}

	bonus := types.Money(float64(card.Balance) * float64(percent) / 100.0 * float64(daysInMonth) / float64(daysInYear))
	if bonus > bonusLimit {
		bonus = bonusLimit
	}

	card.Balance += bonus
}

func Total(cards []types.Card) types.Money {
	var total types.Money = 0
	for _, card := range cards {
		if validateActiveAndBalance(card) {
			total += card.Balance
		}
	}

	return total
}

func PaymentSources(cards []types.Card) []types.PaymentSource {
	paymentSources := []types.PaymentSource{}

	for _, card := range cards {
		if validateActiveAndBalance(card) {
			paymentSources = append(paymentSources, types.PaymentSource{
				Balance: card.Balance,
				Type:    types.TYPE_CARD,
				Number:  card.PAN,
			})
		}
	}

	return paymentSources
}

func validateWithdraw(card types.Card, amount types.Money) bool {
	if !card.Active || card.Balance < amount || amount < 0 || amount > withdrawLimit {
		return false
	}

	return true
}

func validateDeposit(card types.Card, amount types.Money) bool {
	if !card.Active || amount < 0 || amount > depositLimit {
		return false
	}

	return true
}

func validateActiveAndBalance(card types.Card) bool {
	if !card.Active || card.Balance < 0 {
		return false
	}

	return true
}
