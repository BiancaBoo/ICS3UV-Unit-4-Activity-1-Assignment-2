/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2025-11-30
 * @fileoverview This program uses nested loop to print a pattern of numbers.
 */

package main

import "fmt"

func main() {
	var numRows int

	fmt.Print("How many rows would you like? ")
	fmt.Scan(&numRows)

	for i := 1; i <= numRows; i++ {
		rowString := ""
		for j := 1; j <= i; j++ {
			rowString += fmt.Sprintf("%d ", j)
		}
		fmt.Println(rowString)
	}

	fmt.Println("Done.")
}