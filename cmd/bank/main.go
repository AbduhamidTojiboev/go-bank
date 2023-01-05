package main

import (
	"bank/pkg/bank/card"
	"bank/pkg/bank/types"
)

func main() {

	card.IssueCard(types.TJS, "red", "Ivan")

}
