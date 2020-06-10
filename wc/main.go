package main

import (
	"bufio" // read text
	"flag"  // read command line arguments
	"fmt"   // format outout
	"io"    // contains io.Reader and io.Writer interfaces
	"os"    // allows you to use operating system resources
)

func main() {
	// Definging a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")
	// Parsing the flags provided by the user
	flag.Parse()
	fmt.Println(count(os.Stdin, *lines))
}

func count(r io.Reader, countlines bool) int {
	// A scanner is used to read text from the a Reader (such as files)
	scanner := bufio.NewScanner(r)
	// Define scanner split type to words if we are not counting by lines, otherwise default to splitting by lines
	if !countlines {
		scanner.Split(bufio.ScanWords)
	}
	// Defining a counter
	wc := 0

	// For every word scanned, increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}
