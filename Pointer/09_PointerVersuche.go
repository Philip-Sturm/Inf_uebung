package pointer

import "fmt"

func PointerVersuche(x int) {
	y := &x
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(&x)
}
