# go

Key learnings from writing go code.

---

### Go Modules
Go uses file 'go.mod' to track dependencies within the module. The file is put in the root directory and the scope is everything within that folder and subfolders.
- Tracks dependencies just like package.json
- Does not contain any metadata or scripts
- Uses 'go.sum' to lcok the versions of the dependencies

'package main' is the entry point for the application. You cannot directly run a '.go' file with another package name.

When running a ´.go´ file, go will look for 'go.mod' file in the folder and then traverse uppwards until it finds it.

Run 'go help' to se available commands with explainations.

### Pointers
Go is using pointers like C. When passing large objects into functions you can avoid copying it by instead send a reference to the memory adress it is located at. This does not mean that Java always copy over the value into the function. If it is an object then a reference to that object is sent. Java does not allow pointers but use references instead.


Here is an example of using pointers in go.
'''go
package main

import "fmt"

func main() {
	// Declare a variable and a pointer to it
	x := 10
	ptr := &x  // ptr holds the address of x

	// Print the value of x and the value stored at the pointer
	fmt.Println("Value of x:", x)       // Output: 10
	fmt.Println("Pointer to x:", ptr)    // Output: address of x
	fmt.Println("Value at pointer:", *ptr) // Output: 10

	// Modify the value of x through the pointer
	*ptr = 20

	// Print the updated value of x
	fmt.Println("Updated value of x:", x)  // Output: 20
}
'''
