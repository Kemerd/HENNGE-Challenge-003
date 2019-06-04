package main

// Imports
import (
    "fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

// Global Variables
var N int = 0;
var X int = 0;
var err error;
var out int = 0;
var T string = "0";			// All set to 0 just in case there is no input, so we don't crash with nil values.


// Main function.
func main() {
	// Let's grab our input.
    fmt.Print("Enter N: ")
    fmt.Scanln(&N)
	// Make our own recursion.
	loop()
}

// We aren't allowed to use "for," how cruel.
func loop() {
	if N == 0 {return}
	
	// Grab our array length.
	fmt.Scanln(&X)
	
	// Make our tables with our length.
	tNum := make([]string, X)
	
	// Grab our values and put them into an array.
	in := bufio.NewReader(os.Stdin)
	T, err = in.ReadString('\n')
	tNum = strings.Fields(T)
	
	// Parse the numbers, square, and add.
	add(tNum)
	
	// Output and reset.
	fmt.Print(out)
	out = 0;
	
	N--
	loop()
}

// Another loop, until X is 0.
func add(tNum []string) {
	if X == 0 {return}
	
	// Parse a string to an integer. 
	i, err := strconv.Atoi(tNum[X-1])
	if err != nil {}
	
	// If a number is negative, make it 0, so when we add its' square, it does nothing.
	if (i < 0) {
		i = 0;
	}
	
	// Add to our total!
	out = out + i*i

	X--
	add(tNum)
}

// Very interesting challenge. Funnily enough it took me longest to properly parse the input rather than figuring out how to replace the need for a "for" loop.