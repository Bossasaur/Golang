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

	var x int = 10
	var y int = 20

	// and &&
	fmt.Println((x < 100) && (x < 200)) // true
	fmt.Println((x < 300) && (x < 0))   // false

	// or ||
	fmt.Println((x < 0) || (x < 200)) // true
	fmt.Println((x < 0) || (x > 200)) // false

	// not !
	fmt.Println(!(x > y)) // true
	fmt.Println(!(true))  // false
	fmt.Println(!(false)) // true

	// +=
	x += y
	fmt.Println(x) // 30

	// -=
	var xx int = 10
	var yy int = 20

	xx -= yy
	fmt.Println(xx) // -10

	// *=
	var xy int = 10
	var yx int = 20

	xy *= yx
	println(xy) // 200

	// /=
	var xxy int = 200
	var yxx int = 20

	xxy /= yxx
	fmt.Println(xxy) // 10

	// %=
	var xyx int = 210
	var yxy int = 20

	xyx %= yxy
	fmt.Println(xyx) // 10

	// bitwise

	// and &
	var bitx, bity int = 12, 25
	bitz := bitx & bity
	fmt.Println(bitz) // 8

	// or |
	bitz = bitx | bity
	fmt.Println(bitz) // 29

	// XOR ^
	bitz = bitx ^ bity
	fmt.Println(bitz) // 21

	// left shift <<
	bitx = 212
	bitz = bitx << 1
	fmt.Println(bitz) //424

	// right shift >>
	bitz = bitx >> 2
	fmt.Println(bitz) //53

	// control flow

	// if
	var con_a string = "happy"
	if con_a == "happy" {
		fmt.Println(con_a) // happy
	}

	// else
	var fruit string = "grapes"
	if fruit == "apples" {
		fmt.Println("Fruit is an apple")
	} else {
		fmt.Println("Fruit is not an apple") // Fruit is not an Apple
	}

	// if else
	fruit = "oranges"
	if fruit == "apples" {
		fmt.Println("I love apples")
	} else if fruit == "oranges" {
		fmt.Println("Oranges are not an apples") // Oranges are not apples
	} else {
		fmt.Println("No Appetite")
	}

	// switches

	var switch_i int = 100
	switch switch_i {
	case 10:
		fmt.Println("switch_i is 10")
	case 100, 200:
		fmt.Println("switch_i is either 100 or 200") // switch_i is either 100 or 200
	default:
		fmt.Println("switch_i is neither 0, 100 or 200")
	}

	switch_i = 800
	switch switch_i {
	case 10:
		fmt.Println("switch_i is 10")
	case 100, 200:
		fmt.Println("switch_i is either 100 or 200")
	default:
		fmt.Println("switch_i is neither 0, 100 or 200") // switch_i is neither 0, 100 or 200
	}

	// fallthrough
	switch_i = 10
	switch switch_i {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10") // 10
		fallthrough
	case 20:
		fmt.Println("20") // 20
		fallthrough
	default:
		fmt.Println("default") // default
	}

	// conditions within cases
	var switch_x, switch_y int = 10, 20
	switch {
	case switch_x+switch_y == 30:
		fmt.Println("30") // 30
	case switch_x <= switch_y:
		fmt.Println("less than or equal to 30")
	default:
		fmt.Println("greater than 30")
	}

}
