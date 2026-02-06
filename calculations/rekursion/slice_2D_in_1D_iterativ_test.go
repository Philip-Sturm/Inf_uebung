package rekursion

import "fmt"

func ExampleFlattenIterativ() {
	fmt.Println(FlattenIterativ([][]int{
		{1, 2},
		{},
		{3, 4, 5},
		{6},
	}))
	fmt.Println(FlattenIterativ([][]int{}))
	fmt.Println(FlattenIterativ([][]int{{}}))
	fmt.Println(FlattenIterativ([][]int{{42}}))

	// Output:
	// [1 2 3 4 5 6]
	// []
	// []
	// [42]
}
