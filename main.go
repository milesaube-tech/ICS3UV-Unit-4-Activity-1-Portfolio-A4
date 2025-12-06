/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-06
 * @fileoverview This program asks how many integers the user will enter,
 */

package main

import (
	"fmt"
)

func main() {

	// get user input
	var howMany int
	fmt.Print("How many integers will be added: ")
	fmt.Scan(&howMany)

	// set variables
	count := 0
	value := 0
	sum := 0

	// loop to read each integer and add to sum
	for count = 1; count <= howMany; count = count + 1 {
		fmt.Print("Enter an integer: ")
		fmt.Scan(&value)
		sum = sum + value
	}

	// display results
	fmt.Println("The sum is", sum)
}