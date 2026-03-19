package recursions

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountOdd.
MAX. PUNKTE: 10
ZUSATZBEDINGUNG: Die Funktion muss rekursiv sein.
*/

// CountOdd erwartet eine Liste von Zahlen und liefert die Anzahl der ungeraden Zahlen darin.
func CountOdd(list []int) int {
	counter := 0
	if len(list) == 0 {
		return counter
	}
	if list[0]%2 == 1 {
		counter++
		return counter + CountOdd(list[1:])
	}
	return counter + CountOdd(list[1:])
}
