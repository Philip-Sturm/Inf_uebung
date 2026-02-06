package structs

import "fmt"

// Struct zur Darstellung einer komplexen Zahl in Komponentenform
type ComplexNumber struct {
	Real float64
	Imag float64
}

// Gibt die komplexe Zahl als String zurück, z. B. "3 + 2i" oder "4 - 1i"
func (c ComplexNumber) String() string {
	sign := "+"
	if c.Imag < 0 {
		sign = "-"
	}
	return fmt.Sprintf("%.0f %s %.0fi", c.Real, sign, abs(c.Imag))
}

// Hilfsfunktion zur Rückgabe des Betrags eines float64
func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

// AddComplex addiert zwei komplexe Zahlen (a + bi) + (c + di)
func AddComplex(a, b ComplexNumber) ComplexNumber {	
	return ComplexNumber{
		Real: a.Real+b.Real,
		Imag: a.Imag+b.Imag,
	}
}

// MultiplyComplex multipliziert zwei komplexe Zahlen (a + bi) * (c + di)
// Formel: (ac - bd) + (ad + bc)i
func MultiplyComplex(a, b ComplexNumber) ComplexNumber {
	// TODO: Implementiere die Multiplikation von a und b
	return ComplexNumber{
		Real: a.Real*b.Real-a.Imag*b.Imag,
		Imag: a.Real*b.Imag+a.Imag*b.Real,
	}
}
