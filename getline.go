package main

import (
	"bufio"
)

//getLine: get the line wrote by the client
func getLine(scanner *bufio.Scanner) string {
	// Scans a line from Stdin(Console)
	scanner.Scan()
	// Holds the scanned string
	line := scanner.Text()
	return line
}
