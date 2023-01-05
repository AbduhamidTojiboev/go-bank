package main

import (
	"github.com/AbduhamidTojiboev/go-bank/pkg/bank/card"
	"github.com/AbduhamidTojiboev/go-bank/pkg/bank/types"
)

func main() {

	card.IssueCard(types.TJS, "red", "Ivan")

}
