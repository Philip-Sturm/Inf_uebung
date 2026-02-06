package rekursion

import "fmt"

func ExampleSumSlice() {
	fmt.Println(SumSlice([]int{1, 2, 3, 4}))
	fmt.Println(SumSlice([]int{5, -2, 7}))
	fmt.Println(SumSlice([]int{}))
	fmt.Println(SumSlice([]int{10}))

	// Output:
	// 10
	// 10
	// 0
	// 10
}
