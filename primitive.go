package main

import "fmt"

func runPrimitive() {
	var b bool
	fmt.Printf("%v, %T \n", b, b)

	// ! uint8 byte
	var i int32 // ! rune
	fmt.Printf("%v, %T \n", i, i)

	var ui uint32
	fmt.Printf("%v, %T \n", ui, ui)

	x := 10             // ! 1010		2^3
	y := 3              // ! 0011
	fmt.Println(x & y)  // ! 0010
	fmt.Println(x | y)  // ! 1011
	fmt.Println(x ^ y)  // ! 1001
	fmt.Println(x &^ y) // ! 0100

	a := 8
	fmt.Println(a << 3) // ! 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // ! 2^3 / 2^3 = 2^0

	n := 3.14
	n = 3.7e72
	n = 2.1E14
	fmt.Printf("%v, %T \n", n, n)

	var c complex64 = 1 + 2i
	fmt.Printf("%v, %T \n", real(c), real(c))
	fmt.Printf("%v, %T \n", imag(c), imag(c))

	s := "This is a string" // ! immutable
	fmt.Printf("%v, %T \n", s, s)
	fmt.Printf("%v, %T \n", string(s[0]), s[0])

	var r rune = 'a'
	fmt.Printf("%v, %T \n", r, r)
}
