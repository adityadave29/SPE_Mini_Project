package main

import (
	"fmt"
	"math"
)

func validateRangeFloat(number float64) error {
	if number < 0 || number > 9 {
		return fmt.Errorf("enter valid number in range (0-9)")
	}
	return nil
}

func validateRangeInt(number int) error {
	if number < 0 || number > 9 {
		return fmt.Errorf("enter valid number in range (0-9)")
	}
	return nil
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if math.Abs(b) < 1e-9 {
		return 0, fmt.Errorf("number is not divisible by 0")
	}
	return a / b, nil
}

func squareRoot(n float64) (float64, error) {
	if n < 0 {
		return 0, fmt.Errorf("can't take square root of negative")
	}
	return math.Sqrt(n), nil
}

func fact(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial of negative number is not defined")
	}
	if n > 20 {
		return 0, fmt.Errorf("number too large, may cause overflow")
	}

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result, nil
}

func naturalLog(n float64) (float64, error) {
	if n <= 0 {
		return 0, fmt.Errorf("natural log undefined for zero or negative numbers")
	}
	return math.Log(n), nil
}

func pow(a, b float64) float64 {
	return math.Pow(a, b)
}