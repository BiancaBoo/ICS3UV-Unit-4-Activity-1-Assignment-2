// Author: Bianca Boo
// Version 1.0.0
// Date 2025-11-30
// fileoverview: This program uses nested loop to print a pattern of numbers.

package main

import "fmt"

func main() {
	var numRows int

	// number of rows
	fmt.Print("How many rows would you like? ")
	fmt.Scan(&numRows)

	// outer loop
	for i := 1; i <= numRows; i++ {
		rowString := ""

		// inner loop
		for j := 1; j <= i; j++ {
			rowString += fmt.Sprintf("%d ", j)
		}

		// output
		fmt.Println(rowString)
	}

	fmt.Println("Done.")
}