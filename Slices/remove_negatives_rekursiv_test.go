package slices

import "fmt"

func ExampleRemoveNegativesRekursiv() {
	nums := []int{3, -1, 7, -5, 0, 4}

	nums = RemoveNegativesRekursiv(nums)

	for _, n := range nums {
		fmt.Println(n)
	}

	// Output:
	// 3
	// 7
	// 0
	// 4
}
