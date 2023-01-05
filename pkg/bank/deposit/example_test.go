package deposit

import "fmt"

func ExampleCalculateMonth() {
	fmt.Println(CalculateMonth(0, TJS))
	fmt.Println(CalculateMonth(0, USD))
	fmt.Println(CalculateMonth(500_000, TJS))
	fmt.Println(CalculateMonth(500_000, USD))
	fmt.Println(CalculateMonth(1_000_000, TJS))
	fmt.Println(CalculateMonth(1_000_000, USD))
	//Output:
	//0 0
	//0 0
	//1666 2500
	//1250 1666
	//3333 5000
	//2500 3333
}
