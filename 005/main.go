package main

/*
Create a separate file (module) which has at least two methods:

- ReadString: to read a string from console input
- PrintString: to return the string in upper case.

Also create a `main.go` file that acts as calling class.
*/

import (
	"fmt"
)

func main() {
	want := "HELLO WORLD!"

	ReadString()
	got := PrintString()
	// got = "afdsas"

	if got != want {
		fmt.Println(fmt.Errorf("Error: Ex005() = %v, want %v", got, want))
		return
	}

	fmt.Println(got)
}
