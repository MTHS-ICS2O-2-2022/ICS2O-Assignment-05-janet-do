// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file calculates factorials
package main

import (
	"fmt"
	"strconv"
)

func calculate() {
	// Get the input value from the user
	var input string
	fmt.Print("Enter a positive integer: ")
	fmt.Scanln(&input)

	// Parse the input value as an integer
	number, err := strconv.Atoi(input)
	if err != nil || number < 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	// Calculate the factorial using a loop and store the intermediate results
	factorial := 1
	factorialSteps := "1"
	for counter := 2; counter <= number; counter++ {
		factorial *= counter
		factorialSteps += " × " + strconv.Itoa(counter)
	}

	// Display the factorial result along with the intermediate steps
	fmt.Printf("The factorial of %d! is: %d\n", number, factorial)
	fmt.Printf("= %s\n", factorialSteps)
}

func main() {
	calculate()
}
