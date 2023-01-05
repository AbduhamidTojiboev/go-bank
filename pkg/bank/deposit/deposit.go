package deposit

const USD = "USD"
const TJS = "TJS"

func CalucalateYear(amount int, currency string) (min int, max int) {
	var minPercent, maxPercent = PercentFor(currency)
	min = amount * minPercent / 100
	max = amount * maxPercent / 100

	return min, max
}

func PercentFor(currency string) (min int, max int) {
	if currency == USD {
		return 3, 4
	} else if currency == TJS {
		return 4, 6
	}

	return 0, 0
}

func CalculateMonth(amount int, currency string) (min int, max int) {
	min, max = CalucalateYear(amount, currency)
	min /= 12
	max /= 12

	return min, max
}
