package main

import (
	"fmt"
)

func main() {
	// BOOLEAN
	a := true
	fmt.Printf("%v, %T\n", a, a)

	var b bool = true
	fmt.Printf("%v, %T\n", b, b)

	c := 3.14 == 3.14
	d := 3.14 == 14
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)

	// NUMERIC TYPES
	var e uint16 = 14
	fmt.Printf("%v, %T\n", e, e)

	f := 10 // 1010
	g := 3 // 0011
	fmt.Println(f + g)
	fmt.Println(f - g)
	fmt.Println(f * g)
	fmt.Println(f / g)
	fmt.Println(f % g)
	// Bit Operations
	fmt.Println(f & g) // 0010 = 2
	fmt.Println(f | g) // 1011 = 11
	fmt.Println(f ^ g) // exclusive OR 1001
	fmt.Println(f &^ g) // in NOT 0100
	// Bit Shifting
	h := 8 // 2^3
	fmt.Println(h << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(h >> 3) // 2^3 / 2^3 = 2^0
	// Types can't change when operations are done
	// int / int will produce an int
}
