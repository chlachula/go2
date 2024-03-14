package pi

import (
	"fmt"
	"math"
)

// This function calculates PI using the Leibniz formula
func CalculatePi_byLeibniz(iterations int) float64 {
	var pi float64
	sign := 1.0

	for i := 0; i < iterations; i++ {
		term := 1.0 / (float64(2*i) + 1)
		pi += sign * term
		sign *= -1.0 // Alternate signs for each term
	}

	return 4 * pi
}

func Leibniz() {
	var iterations int
	/*
	   fmt.Println("Enter desired precision (number of terms):")
	   fmt.Scanf("%d", &precision)
	*/
	iterations = 2100100100
	if iterations <= 0 {
		fmt.Println("Error: Precision must be a positive integer.")
		return
	}

	pi := CalculatePi_byLeibniz(iterations)
	fmt.Printf("PI (using %d terms): %.10f\n", iterations, pi)

	// Optional: Compare with math.Pi constant
	fmt.Printf("Difference from math.Pi: %.10f\n", math.Abs(pi-math.Pi))
}

func BaileyBorweinPlouffe() {

}
