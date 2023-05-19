// Copyright (c) 2020 Janet Do All rights reserved
//
// Created by: Janet Do
// Created on: Sep 2020
// This file calculates factorials

"use strict"

function calculate() {
  // Get the input value from the user
  const input = document.getElementById("integer").value

  // Parse the input value as an integer
  const number = parseInt(input)

  // Check if the input is a positive integer
  if (isNaN(number) || number < 0) {
    document.getElementById("answer").innerHTML = "<p>Please enter a positive integer.</p>"
    return
  }

  // Calculate the factorial using a loop and store the intermediate results
  let factorial = 1
  let factorialSteps = "1"
  for (let counter = 2; counter <= number; counter++) {
    factorial *= counter
    factorialSteps += " × " + counter
  }

  // Display the factorial result along with the intermediate steps
  document.getElementById("answer").innerHTML =
    "<p>The factorial of " + number + "! is: " + factorial + "</p>" +
    "<p>4!: " + factorialSteps + "</p>"
}
