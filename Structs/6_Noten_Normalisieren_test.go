package structs

import "fmt"

func ExampleNormalizeGrades() {
	students := []Student{
		{"Anna", 85},
		{"Ben", 102},
		{"Clara", -5},
	}

	NormalizeGrades(students)

	for _, s := range students {
		fmt.Printf("%s: %d\n", s.Name, s.Grade)
	}

	// Output:
	// Anna: 85
	// Ben: 100
	// Clara: 0
}
