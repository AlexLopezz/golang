package main // <-- All go files must be a package.

import "fmt" // <-- or when we used more than 1: import ( "package" "package" "package" )

func main() { // entrypoint on go app.
	fmt.Println("Hello World, I'm Alex :)") // print in console.
}

/*
This is your Go code. In this code, you:
	* Declare a main package (a package is a way to group functions, and it's made up of all the files in the same directory).
	* Import the popular fmt package, which contains functions for formatting text, including printing to the console. This package is one of the standard library packages you got when you installed Go.
	* Implement a main function to print a message to the console. A main function executes by default when you run the main package.
*/

// For execute this file: $ go run hello-world.go
