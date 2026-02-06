package slices

import "fmt"

func ExampleRemoveShortWords() {
	words := []string{"Go", "ist", "super", "!", "Hi", "ChatGPT"}

	words = RemoveShortWords(words, 3)

	for _, w := range words {
		fmt.Println(w)
	}

	// Output:
	// ist
	// super
	// ChatGPT
}
