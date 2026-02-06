package rekursion

import "fmt"

func ExampleDigitSum() {
	fmt.Println(DigitSum(1234))
	fmt.Println(DigitSum(9))
	fmt.Println(DigitSum(5050))
	fmt.Println(DigitSum(0))

	// Output:
	// 10
	// 9
	// 10
	// 0
}
