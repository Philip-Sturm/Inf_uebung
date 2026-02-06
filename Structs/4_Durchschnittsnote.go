package structs

import "fmt"

// Struct zur Darstellung einer Bewertung
type Rating struct {
	Name  string // Name der bewertenden Person
	Score int    // Bewertung zwischen z. B. 1 und 5
}

// AverageScore berechnet den Durchschnittswert aller Bewertungen
func AverageScore(ratings []Rating) float64 {
	// TODO: Implementieren
	var Total int
	var average float64
	if len(ratings)==0 {
		return 0.0
	}
	fmt.Println(Total)

	return average // Platzhalter
}
