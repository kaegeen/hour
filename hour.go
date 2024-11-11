package main

import (
	"fmt"
)

// Function to calculate hours in a given number of years
func hoursInYears(years int) int {
	// Average number of hours in a year (ignoring leap years for simplicity)
	hoursPerYear := 365 * 24
	// Calculating total hours
	return years * hoursPerYear
}

func main() {
	// Ask the user for input (number of years)
	var years int
	fmt.Println("Enter the number of years to calculate hours:")
	fmt.Scan(&years)

	// Get the total number of hours for the input years
	totalHours := hoursInYears(years)

	// Output the result
	fmt.Printf("There are %d hours in %d year(s).\n", totalHours, years)

	// For additional examples:
	fmt.Println("Additional calculations:")
	fmt.Printf("There are %d hours in 10 years.\n", hoursInYears(10))
	fmt.Printf("There are %d hours in 100 years.\n", hoursInYears(100))
	fmt.Printf("There are %d hours in 1000 years.\n", hoursInYears(1000))
}
