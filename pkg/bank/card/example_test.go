package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleWithdraw_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 1000000
}

func ExampleWithdraw_noMoney() {
	card := types.Card{Balance: 0, Active: true}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 0
}

func ExampleWithdraw_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Withdraw(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleWithdraw_limit() {
	card := types.Card{Balance: 200_000_00, Active: true}
	Withdraw(&card, 30_000_00)
	fmt.Println(card.Balance)
	// Output: 20000000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 3000000
}

func ExampleDeposit_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 200_000_00, Active: true}
	Deposit(&card, 60_000_00)
	fmt.Println(card.Balance)
	// Output: 20000000
}

func ExampleAddBonus_positive() {
	card := types.Card{Balance: 10_000_00, Active: true}
	AddBonus(&card, PercentBonus, 30, 365)
	fmt.Println(card.Balance)
	// Output: 1002465
}

func ExampleAddBonus_noMoney() {
	card := types.Card{Balance: -10, Active: true}
	AddBonus(&card, PercentBonus, 30, 365)
	fmt.Println(card.Balance)
	// Output: -10
}

func ExampleAddBonus_inactive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	AddBonus(&card, PercentBonus, 30, 365)
	fmt.Println(card.Balance)
	// Output: 2000000
}

func ExampleAddBonus_limit() {
	card := types.Card{Balance: 3_000_000_00, Active: true}
	AddBonus(&card, PercentBonus, 30, 365)
	fmt.Println(card.Balance)
	// Output: 300500000
}

func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
		},
		{
			Balance: 10_000_00,
			Active:  false,
		},
		{
			Balance: -10_000_00,
			Active:  false,
		},
	}
	fmt.Println(Total(cards))
	//Output: 1000000
}

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active:  true,
			PAN:     "12121 21212 1212",
		},
		{
			Balance: 50_000_00,
			Active:  true,
			PAN:     "12121 12 1212",
		},
		{
			Balance: -10_000_00,
			Active:  false,
			PAN:     "12121 21212 1212",
		},
	}
	paymentSources := PaymentSources(cards)

	fmt.Println(len(paymentSources))
	//Output: 2
}
