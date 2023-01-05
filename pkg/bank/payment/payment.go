package payment

import (
	"github.com/AbduhamidTojiboev/go-bank/pkg/bank/types"
)

func Max(payments []types.Payment) types.Payment {
	if len(payments) == 0 {
		return types.Payment{}
	}

	maxIndex := 0
	for i, payment := range payments {
		if payment.Amount > payments[maxIndex].Amount {
			maxIndex = i
		}
	}
	return payments[maxIndex]
}
