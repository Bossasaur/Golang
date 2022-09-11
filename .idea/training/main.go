package main

import "fmt"

// comment

/*
 multi
 line
 comment
*/

func main() {
	var place string = "Mindera"
	var firstname string = "Susan"
	var lastname string = "Lockwood"
	var space string = " "
	message := "\n" + "Hello" + space + firstname + space + lastname + space + "from" + space + place + "\n\n"
	fmt.Print(message)
}
