package slices

import "fmt"

func ExampleCutAfterWord() {
	words := []string{"Go", "is", "a", "fun", "and", "powerful", "language"}

	words = CutAfterWord(words, "fun")

	for _, w := range words {
		fmt.Println(w)
	}

	// Output:
	// Go
	// is
	// a
	// fun
}
