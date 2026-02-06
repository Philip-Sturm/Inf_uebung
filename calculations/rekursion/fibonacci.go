package rekursion

// Erwartet eine ganze Zahl n >= 0.
// Liefert die n-te Zahl der Fibonacci-Folge.
// Die Funktion soll rekursiv implementiert werden.
//
// Beispiele für die Folge:
// n = 0 → 0
// n = 1 → 1
// n = 2 → 1
// n = 3 → 2
// n = 4 → 3
// n = 5 → 5
//
// Hinweis: Die rekursive Definition lautet:
// Fibonacci(n) = Fibonacci(n-1) + Fibonacci(n-2)
// mit den Basisfällen: Fibonacci(0) = 0, Fibonacci(1) = 1
func Fibonacci2(n int) int {
	if n<=1 {
		return n
	}
	return Fibonacci2(n-1)+Fibonacci2(n-2)
}

func Fibonacci(n int) int {
	var fibn1 = 1
	var fibn2 = 0
	var next int
	
	if n==0 {return fibn2}
	if n==1 {return fibn1}
	

	for i := 1; i < n; i++ {
		next = fibn1 + fibn2
		fibn2 = fibn1
		fibn1 = next
	}
	return next
}