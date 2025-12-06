/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-06
 * @fileoverview This program asks how many integers the user will enter, reads that many
 */

// get user input
let howMany: number = Number(prompt("How many integers will be added:"));

// set variables
let count: number = 0;
let value: number = 0;
let sum: number = 0;

// loop to read each integer and add to sum
for (count = 1; count <= howMany; count = count + 1) {
  value = Number(prompt("Enter an integer:"));
  sum = sum + value;
}

// display results
console.log("The sum is " + sum);