package rekursion

import "fmt"

func ExampleFlattenRekursiv() {
	fmt.Println(FlattenRekursiv([][]int{
		{1, 2},
		{},
		{3, 4, 5},
		{6},
	}))
	fmt.Println(FlattenRekursiv([][]int{}))
	fmt.Println(FlattenRekursiv([][]int{{}}))
	fmt.Println(FlattenRekursiv([][]int{{42}}))

	// Output:
	// [1 2 3 4 5 6]
	// []
	// []
	// [42]
}
