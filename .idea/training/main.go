package main

import "fmt"

// comment

/*
 multi
 line
 comment
*/

func main() {
	place := "Mindera"
	firstname := "Susan"
	lastname := "Lockwood"
	space := " "
	message := "Hello " + firstname + space + lastname + space + "from" + space + place
	fmt.Println(message)
}
