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

// https://en.wikipedia.org/wiki/Bailey%E2%80%93Borwein%E2%80%93Plouffe_formula
// https://go.dev/play/p/WPABH3SPdRu
func BaileyBorweinPlouffe(kmax int) {
	var sum, k8, k16 float64
	k16 = 1.
	for k := 0; k <= kmax; k++ {
		k8 = float64(k) * 8.0
		sum += (k16 * (4./(k8+1.) - 2./(k8+4.) - 1./(k8+5.) - 1./(k8+6.)))
		k16 /= 16.0
	}
	fmt.Printf("PI (using BBP formula, %d digits): %.10f\n", kmax, sum)
	/*	var (
			precision int = 30 // Adjust for desired precision (number of decimal places)
			k         *big.Int
			piStr     string
		)

		// Initialize big.Int variables with high precision
		k = new(big.Int)
		k.SetString("16", 10) // Base 16 for BBP formula

		// Define constants used in the formula
		one := new(big.Int)
		one.SetInt64(1)
		four := new(big.Int)
		four.SetInt64(4)
		five := new(big.Int).SetInt64(5)
		six := new(big.Int).SetInt64(6)
		eight := new(big.Int).SetInt64(8)
		ten := new(big.Int).SetInt64(10)

		// Perform BBP formula calculation iteratively
		for i := 0; i < precision; i++ {
			term1 := new(big.Int)
			term1.Mul(k, k)
			term1.Add(term1, one)

			term2 := new(big.Int)
			term2.Mul(k, k)
			term2.Add(term2, four)

			term3 := new(big.Int)
			term3.Mul(k, k)
			term3.Add(term3, five)

			term4 := new(big.Int)
			term4.Mul(k, k)
			term4.Add(term4, six)

			// Calculate the summation and update piStr
			sum := new(big.Int)
			sum.Sub(term1, term2)
			sum.Sub(sum, term3)
			sum.Sub(sum, term4)
			piStr += fmt.Sprintf("%d", sum.Mod(sum, ten))

			// Increment k for the next iteration
			k.Add(k, eight)
		}

		// Convert the calculated digits to a float64 value
		//	pi, _ := new(big.Float).SetString(piStr, 10)
		pi, _ := new(big.Float).SetString(piStr)

		fmt.Printf("PI (using BBP formula, %d digits): %.10f\n", precision, pi)
	*/
}
