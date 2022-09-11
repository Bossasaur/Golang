package main

import "fmt"

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
	// variable scoping
	{
		state := "PA"
		fmt.Println(city)
		fmt.Println(state)
		fmt.Println("")
		fmt.Println("")
	}
}
