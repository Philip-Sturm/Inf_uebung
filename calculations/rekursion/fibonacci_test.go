package rekursion

import (
	"fmt"
)

func ExampleFibonacci() {
	fmt.Println(Fibonacci(0))
	fmt.Println(Fibonacci(1))
	fmt.Println(Fibonacci(2))
	fmt.Println(Fibonacci(5))
	fmt.Println(Fibonacci(10))


	// Output:
	// 0
	// 1
	// 1
	// 5
	// 55
	
}
