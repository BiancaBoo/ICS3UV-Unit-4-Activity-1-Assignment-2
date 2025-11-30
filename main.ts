/**
 * @author Bianca Boo
 * @version 1.0.0
 * @date 2025-11-30
 * @fileoverview This program uses nested loop to print a pattern of numbers.
 */

// variables
let numRowsString: string = "";
let numRows: number = 0;

// number of rows
numRowsString = prompt("How many rows would you like? ") || "0";
numRows = parseInt(numRowsString);

// outer loop
for (let i: number = 1; i <= numRows; i++) {
  let rowString: string = "";

  // inner loop
  for (let j: number = 1; j <= i; j++) {
    rowString += j + " ";
  }

  // output
  console.log(rowString.trim()); 
}

console.log("Done.");