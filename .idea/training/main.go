package main

import "fmt"

// global variable
const country string = "United States"

func main() {
	var loc string
	var fn string
	var ln string
	var yr int
	loc, fn, ln, yr = "Mindera", "Susan", "Lockwood", 2022
	fmt.Println("")
	fmt.Printf("Hello %v %v from %v! Welcome to %d!", fn, ln, loc, yr)
	fmt.Println("")
	fmt.Println("")
	name := "Daniel Nations" // short variable declaration
	fmt.Printf("Hello %v from %v! Welcome to %d!", name, loc, yr)
	fmt.Println("")
	fmt.Println("")
	city := "Pittsburgh"
	// variable scoping - inner block
	{
		state := "PA"
		fmt.Printf("This is %v, %v, %v", city, state, country)
		fmt.Println("")
		fmt.Println("")
	}
	// operators
	var this_city string = "Pittsburgh"
	var this_city_2 string = "Ellwood"
	fmt.Println(this_city == this_city_2) // false
	fmt.Println(this_city != this_city_2) // true

	var a, b int = 5, 10
	fmt.Println(a < b) // true

	var d, c int = 10, 10
	fmt.Println(d <= c) // true

	var e, f int = 20, 10
	fmt.Println(e > f) // true

	var g, h int = 20, 20
	fmt.Println(g >= h) // true

	//increments
	var inc int = 1
	inc++
	fmt.Println(inc) //2

	//decrement
	var dec int = 2
	dec--
	fmt.Println(dec) // 1

}
