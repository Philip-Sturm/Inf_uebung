package recursions

//Die Funktion factorial soll das factorial recursiv errechen der gegebenen Zahl
func Factorial(Num int) int {
	if Num == 0 {
		return 1
	}
	return Num * Factorial(Num-1)
}
