package rekursion

// Erwartet einen Slice von ganzen Zahlen (ints).
// Gibt die Summe aller Zahlen im Slice zurück.
// Die Funktion soll rekursiv implementiert werden.
//
// Beispiel:
// SumSlice([]int{1, 2, 3, 4}) → 10
// SumSlice([]int{5, -2, 7})  → 10
// SumSlice([]int{})          → 0
//
func SumSlice(numbers []int) int {
	return SumSliceWithCounter(numbers, 0)	
}

func SumSliceWithCounter(numbers []int, counter int) int {
	var newcounter int = counter
	newcounter++
	if counter >= len(numbers) {
		return 0
	}
	return SumSliceWithCounter(numbers, newcounter)+numbers[counter]
}