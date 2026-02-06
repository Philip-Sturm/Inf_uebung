package rekursion

import "fmt"

func ExampleCountLetterNested() {
	input := [][]string{
		{"Apfel", "Banane"},
		{"Ananas"},
		{},
		{"Avocado", "Apfel"},
	}
	fmt.Println(CountLetterNested(input, 'a'))
	fmt.Println(CountLetterNested(input, 'n'))
	fmt.Println(CountLetterNested(input, 'x'))

	// Output:
	// 7
	// 3
	// 0
}
