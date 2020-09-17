package main

import "fmt"

func runIterable() {
	grades := [...]int{98, 97, 99}
	grades[1] = 100
	fmt.Printf("Grades: %v, %T", grades, len(grades))

	garades := &grades
	fmt.Printf("Garades: %v", garades)

	// ! Slices
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a[:]
	c := a[3:]
	d := a[:3]
	fmt.Printf("Slices %v %v %v %v", a, b, c, d)

	a = append(a, 10)
	fmt.Printf("Slices now %v", a)
}
