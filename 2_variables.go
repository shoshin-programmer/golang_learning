package main

import (
	// commas are not use, instead new lines are used for the list of packages
	"fmt"
	"strconv" // str conversion packages
)

// VARIABLES
var d float32 = 4

var (
	character string = "Rick"
	companion string = "Morty"
)

func main() {
	// METHOD 1
	var a int
	a = 1
	fmt.Println(a)
	// useful when getting a value later

	// METHOD 2
	var b int = 2
	fmt.Println(b)
	// verbose enough

	// METHOD 3
	c := 3
	fmt.Println(c)
	// concise and can also be used inside the main

	// Outside variable int
	fmt.Printf("%v, %T\n", d, d)

	fmt.Printf("COERCING TYPES\n")
	// Coercing types
	var e float32
	e = d
	fmt.Printf("%v, %T\n", e, e)

	var f float32 = 3.14
	fmt.Printf("%v, %T\n", f, f)
	var g int
	g = int(f)
	// g = f THIS WON'T WORK
	// From float to int, the float decimal value was lost
	fmt.Printf("%v, %T\n", g, g)

	var h string
	h = string(g)
	// str is looking for the unicode character with value of the int
	fmt.Printf("%v, %T\n", h, h)

	var i string
	i = strconv.Itoa(g)
	// used to convert int to string
	fmt.Printf("%v, %T\n", i, i)

	j := "test string"
	fmt.Printf("%v, %T\n", j, j)

	// Outside variable
	fmt.Printf("%v, %v\n", character, companion)

	// NOTES
	// test := 21
	// Above can't be used if there is already a variable named test
	// Capitalized first character variables are of global scope

}
