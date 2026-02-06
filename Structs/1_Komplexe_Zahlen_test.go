package structs

import "fmt"

func ExampleAddComplex() {
	var a ComplexNumber = ComplexNumber{Real: 2, Imag: 3}  // 2 + 3i
	var b ComplexNumber = ComplexNumber{Real: 1, Imag: -1} // 1 - 1i

	sum := AddComplex(a, b)
	//product := MultiplyComplex(a, b)

	fmt.Println(sum.String())
	//fmt.Println(product.String())

	// Output:
	// 3 + 2i
}

func ExampleMultiplyComplex() {
	var a ComplexNumber = ComplexNumber{Real: 2, Imag: 3}  // 2 + 3i
	var b ComplexNumber = ComplexNumber{Real: 1, Imag: -1} // 1 - 1i

	//sum := AddComplex(a, b)
	product := MultiplyComplex(a, b)

	//fmt.Println(sum.String())
	fmt.Println(product.String())

	// Output:
	// 5 + 1i
}
