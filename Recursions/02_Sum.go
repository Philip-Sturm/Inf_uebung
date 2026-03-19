package recursions

//Die Funktion Sum soll die Summe recursiv errechen der Zahlen bis Num
func Sum(Num int) int {
	if Num == 0 {
		return 0
	}
	return Num + Sum(Num-1)
}
