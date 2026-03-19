package recursions

import "fmt"

func ExampleFactorial() {
	fmt.Println(Factorial(6))
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(42))

	// Output:
	//
}

func ExampleSum() {
	fmt.Println(Sum(6))
	fmt.Println(Sum(0))
	fmt.Println(Sum(42))

	// Output:
	//
}

func ExampleCountOdd() {
	l1 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	l2 := []int{1, 3, 5, 7, 9}
	l3 := []int{2, 4, 6, 8, 10}
	l4 := []int{}

	fmt.Println(CountOdd(l1))
	fmt.Println(CountOdd(l2))
	fmt.Println(CountOdd(l3))
	fmt.Println(CountOdd(l4))

	// Output:
	// 4
	// 5
	// 0
	// 1
}
