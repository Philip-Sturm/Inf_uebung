package structs

import "fmt"

func ExampleAverageScore() {
	ratings := []Rating{
		{Name: "Anna", Score: 4},
		{Name: "Ben", Score: 5},
		{Name: "Clara", Score: 3},
		{Name: "David", Score: 4},
	}

	avg := AverageScore(ratings)
	fmt.Printf("%.2f\n", avg)

	// Output:
	// 4.00
}
