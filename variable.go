package main

import (
	"fmt"
	"strconv"
)

var (
	k = 42
	l = 52
)

func runVariable() {
	var i int
	i = 42
	fmt.Printf("i is %d \n", i)

	j := 99
	fmt.Printf("j is %v \n", j)

	// package
	fmt.Printf("%v %v \n", k, l)

	m := float32(i)
	fmt.Printf("Konverted %v \n", m)

	n := string(i)
	fmt.Printf("Konverted %v \n", n)

	o := strconv.Itoa(i)
	fmt.Printf("Konverted to %v %T \n", o, o)

}
