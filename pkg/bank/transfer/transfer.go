package transfer

func Total(amount int) int {
	var bonusPercent = 0.5

	return amount + int(float64(amount)*bonusPercent/100)
}
