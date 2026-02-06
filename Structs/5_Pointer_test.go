package structs

import "fmt"

func ExampleConvertToFahrenheit() {
	t1 := 0.0
	t2 := 100.0
	t3 := -40.0

	ConvertToFahrenheit(&t1)
	ConvertToFahrenheit(&t2)
	ConvertToFahrenheit(&t3)

	fmt.Printf("%.1f\n", t1)
	fmt.Printf("%.1f\n", t2)
	fmt.Printf("%.1f\n", t3)

	// Output:
	// 32.0
	// 212.0
	// -40.0
}
