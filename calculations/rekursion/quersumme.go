package rekursion

// Erwartet eine natürliche Zahl n >= 0.
// Gibt die Summe aller Ziffern von n zurück.
// Die Funktion soll rekursiv implementiert werden.
//
// Beispiel:
// DigitSum(1234) → 1 + 2 + 3 + 4 = 10
// DigitSum(9)    → 9
// DigitSum(5050) → 5 + 0 + 5 + 0 = 10
//
// Hinweis: Eine mögliche rekursive Zerlegung:
// DigitSum(n) = letzte Ziffer + DigitSum(alle anderen Ziffern)
// z.B. DigitSum(123) = 3 + DigitSum(12)
func DigitSum(n int) int {
	if(n==0){
		return 0
	}
	return n%10 + DigitSum(n/10)
}
