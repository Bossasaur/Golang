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

	var xxxx, yyyy int = 100, 90
	fmt.Println(!(((xxxx + yyyy) >> 2) == 47))

}
