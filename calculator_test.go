package main

import (
	"math"
	"testing"
)

// adityadave@Adityas-MacBook-Air-3 Project % go test ./...
/*
stage('Build') {
    steps {
        sh 'go mod tidy'
        sh 'go test ./...'
        sh 'go build -o calculator'
    }
}
*/

//============================ Addition Testcases ==================================
func TestAdd1(t *testing.T){
	result := add(9,10)
	if result != 20{
		t.Errorf("Expected 19,got %f",result)
	}
}

func TestAdd2(t *testing.T){
	result := add(-5,4)
	if result != -1{
		t.Errorf("Expected -1, got %f",result)
	}
}

func TestAdd3(t *testing.T){
	result := add(-3,-4)
	if result != -7{
		t.Errorf("Expected -7, got %f", result)
	}
}

func TestAdd4(t *testing.T) {
    result := add(math.MaxFloat64, math.MaxFloat64)
    if !math.IsInf(result, 1) {
        t.Errorf("Expected +Inf, got %g", result)
    }
}

func TestAdd5(t *testing.T) {
    result := add(-math.MaxFloat64, -math.MaxFloat64)
    if !math.IsInf(result, -1) {
        t.Errorf("Expected -Inf, got %g", result)
    }
}

func TestAdd6(t *testing.T) {
    result := add(3,0)
    if result != 3{
        t.Errorf("Expected -Inf, got %g", result)
    }
}

func TestAdd7(t *testing.T) {
    // 1e16 is large enough that adding 1 might be lost in 64-bit floats
    result := add(1e16, 1)
    if result != 10000000000000001 {
        t.Errorf("Expected 10000000000000001, got %g", result)
    }
}

func TestAdd8(t *testing.T) {
    result := add(math.Inf(1), math.Inf(-1))
    if !math.IsNaN(result) {
        t.Errorf("Expected NaN, got %g", result)
    }
}

func TestAdd9(t *testing.T) {
    result := add(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64)
    if result == 0 {
        t.Errorf("Expected non-zero result, got %g", result)
    }
}

// =================================== Subtraction Testcases ===============================================

func TestSubtraction1(t *testing.T) {
    result := subtract(20, 1)
    if result != 19 {
        t.Errorf("Expected 19, got %g", result)
    }
}

func TestSubtraction2(t *testing.T) {
    result := subtract(-5, 4)
    if result != -9 {
        t.Errorf("Expected -9, got %g", result)
    }
}

func TestSubtraction3(t *testing.T) {
    result := subtract(-3, -4)
    if result != 1 {
        t.Errorf("Expected 1, got %g", result)
    }
}

func TestSubtraction4(t *testing.T) {
    // Overflow: Subtracting a huge negative from a huge positive
    result := subtract(math.MaxFloat64, -math.MaxFloat64)
    if !math.IsInf(result, 1) {
        t.Errorf("Expected +Inf, got %g", result)
    }
}

func TestSubtraction5(t *testing.T) {
    // Underflow: Subtracting a huge positive from a huge negative
    result := subtract(-math.MaxFloat64, math.MaxFloat64)
    if !math.IsInf(result, -1) {
        t.Errorf("Expected -Inf, got %g", result)
    }
}

func TestSubtraction6(t *testing.T) {
    result := subtract(3, 0)
    if result != 3 {
        t.Errorf("Expected 3, got %g", result)
    }
}

func TestSubtraction7(t *testing.T) {
    // Precision: Checking if a small change is registered at high scales
    result := subtract(10000000000000001, 1)
    if result != 1e16 {
        t.Errorf("Expected 1e16, got %g", result)
    }
}

func TestSubtraction8(t *testing.T) {
    // Subtracting Infinity from Infinity is undefined (NaN)
    result := subtract(math.Inf(1), math.Inf(1))
    if !math.IsNaN(result) {
        t.Errorf("Expected NaN, got %g", result)
    }
}

func TestSubtraction9(t *testing.T) {
    // Zero result from identical small numbers
    result := subtract(math.SmallestNonzeroFloat64, math.SmallestNonzeroFloat64)
    if result != 0 {
        t.Errorf("Expected 0, got %g", result)
    }
}

//================================== Multiplication Testcases ==================================

func TestMultiplication1(t *testing.T) {
    result := multiply(9, 10)
    if result != 90 {
        t.Errorf("Expected 90, got %g", result)
    }
}

func TestMultiplication2(t *testing.T) {
    // Negative by Positive
    result := multiply(-5, 4)
    if result != -20 {
        t.Errorf("Expected -20, got %g", result)
    }
}

func TestMultiplication3(t *testing.T) {
    // Negative by Negative
    result := multiply(-3, -4)
    if result != 12 {
        t.Errorf("Expected 12, got %g", result)
    }
}

func TestMultiplication4(t *testing.T) {
    // Overflow: Two large numbers exceeding MaxFloat64
    result := multiply(1e200, 1e200)
    if !math.IsInf(result, 1) {
        t.Errorf("Expected +Inf, got %g", result)
    }
}

func TestMultiplication5(t *testing.T) {
    // Zero Identity
    result := multiply(math.MaxFloat64, 0)
    if result != 0 {
        t.Errorf("Expected 0, got %g", result)
    }
}

func TestMultiplication6(t *testing.T) {
    // Sign of Zero: -5 * 0 should ideally be -0 or 0
    result := multiply(-5, 0)
    if result != 0 && result != -0 {
        t.Errorf("Expected 0, got %g", result)
    }
}

func TestMultiplication7(t *testing.T) {
    // Precision: Small decimals
    result := multiply(0.1, 0.2)
    // Note: result might be 0.020000000000000004 due to float logic
    if math.Abs(result-0.02) > 1e-10 {
        t.Errorf("Expected 0.02, got %g", result)
    }
}

func TestMultiplication8(t *testing.T) {
    // Multiplying Infinity by Zero is undefined
    result := multiply(math.Inf(1), 0)
    if !math.IsNaN(result) {
        t.Errorf("Expected NaN, got %g", result)
    }
}

// =========================== Division Testcases ====================================================

func TestDivision1(t *testing.T) {
    // Basic division
    result, err := divide(100, 10)
    if err != nil || result != 10 {
        t.Errorf("Expected 10, got %g (Error: %v)", result, err)
    }
}

func TestDivision2(t *testing.T) {
    // Division resulting in a decimal
    result, err := divide(1, 4)
    if err != nil || result != 0.25 {
        t.Errorf("Expected 0.25, got %g", result)
    }
}

func TestDivision3(t *testing.T) {
    // Negative result
    result, err := divide(-10, 2)
    if err != nil || result != -5 {
        t.Errorf("Expected -5, got %g", result)
    }
}

func TestDivision4(t *testing.T) {
    // Division by Zero (The Hard Zero)
    _, err := divide(10, 0)
    if err == nil {
        t.Errorf("Expected error for division by zero, got nil")
    }
}

func TestDivision5(t *testing.T) {
    // Division by "Near-Zero" (Your 1e-9 threshold)
    // Since 1e-10 < 1e-9, this should trigger your error
    _, err := divide(10, 1e-10)
    if err == nil {
        t.Errorf("Expected error for threshold underflow, got nil")
    }
}

func TestDivision6(t *testing.T) {
    // Dividing Zero by a number
    result, err := divide(0, 5)
    if err != nil || result != 0 {
        t.Errorf("Expected 0, got %g", result)
    }
}

func TestDivision7(t *testing.T) {
    // Resulting in a very small number (Underflow)
    result, err := divide(1, math.MaxFloat64)
    if err != nil || result != 0 {
        // Technically it's a subnormal number, but usually rounds to 0
        if result > 1e-300 { 
             t.Errorf("Expected near-zero, got %g", result)
        }
    }
}

func TestDivision8(t *testing.T) {
    // Dividing Infinity
    result, err := divide(math.Inf(1), 2)
    if err != nil || !math.IsInf(result, 1) {
        t.Errorf("Expected +Inf, got %g", result)
    }
}

func TestDivision9(t *testing.T) {
    // Dividing Infinity by Infinity (Undefined)
    result, err := divide(math.Inf(1), math.Inf(1))
    if err != nil || !math.IsNaN(result) {
        t.Errorf("Expected NaN, got %g", result)
    }
}

// =================================== Square Root Testcases =========================================

func TestSquareRoot1(t *testing.T) {
	// Standard perfect square
	result, err := squareRoot(25)
	if err != nil || result != 5 {
		t.Errorf("Expected 5, got %g (Error: %v)", result, err)
	}
}

func TestSquareRoot2(t *testing.T) {
	// Square root of zero
	result, err := squareRoot(0)
	if err != nil || result != 0 {
		t.Errorf("Expected 0, got %g", result)
	}
}

func TestSquareRoot3(t *testing.T) {
	// Negative number (Should return error)
	_, err := squareRoot(-4)
	if err == nil {
		t.Errorf("Expected error for negative input, got nil")
	}
}

func TestSquareRoot4(t *testing.T) {
	// Decimal/Irrational result
	result, err := squareRoot(2)
	// We check for a close approximation because math.Sqrt(2) isn't exactly representable
	if err != nil || math.Abs(result-1.41421356) > 1e-8 {
		t.Errorf("Expected ~1.4142, got %g", result)
	}
}

func TestSquareRoot5(t *testing.T) {
	// Square root of the maximum float value
	result, err := squareRoot(math.MaxFloat64)
	if err != nil || result <= 0 {
		t.Errorf("Expected a large positive number, got %g", result)
	}
}

func TestSquareRoot6(t *testing.T) {
	// Square root of a very tiny positive number
	result, err := squareRoot(1e-10)
	if err != nil || result != 1e-5 {
		t.Errorf("Expected 0.00001, got %g", result)
	}
}

func TestSquareRoot7(t *testing.T) {
	// Square root of Infinity
	result, err := squareRoot(math.Inf(1))
	if err != nil || !math.IsInf(result, 1) {
		t.Errorf("Expected +Inf, got %g", result)
	}
}

func TestSquareRoot8(t *testing.T) {
	// Square root of a number very close to zero (Subnormal)
	result, err := squareRoot(math.SmallestNonzeroFloat64)
	if err != nil || result == 0 {
		t.Errorf("Expected non-zero result for smallest positive float, got %g", result)
	}
}

// ========================== Factorial Testcases ===============================================

func TestFactorial1(t *testing.T) {
    // Base Case: 0! is mathematically 1
    result, err := fact(0)
    if err != nil || result != 1 {
        t.Errorf("Expected 1, got %d (Error: %v)", result, err)
    }
}

func TestFactorial2(t *testing.T) {
    // Base Case: 1! is 1
    result, err := fact(1)
    if err != nil || result != 1 {
        t.Errorf("Expected 1, got %d", result)
    }
}

func TestFactorial3(t *testing.T) {
    // Standard small integer
    result, err := fact(5)
    if err != nil || result != 120 {
        t.Errorf("Expected 120, got %d", result)
    }
}

func TestFactorial4(t *testing.T) {
    // Negative input (Error handling)
    _, err := fact(-1)
    if err == nil {
        t.Errorf("Expected error for negative input, got nil")
    }
}

func TestFactorial5(t *testing.T) {
    // Boundary Case: Maximum allowed input (20)
    // 20! = 2,432,902,008,176,640,000
    result, err := fact(20)
    if err != nil || result != 2432902008176640000 {
        t.Errorf("Expected 2432902008176640000, got %d", result)
    }
}

func TestFactorial6(t *testing.T) {
    // Boundary Case: Exceeding the max limit (Overflow prevention)
    _, err := fact(21)
    if err == nil {
        t.Errorf("Expected overflow error for input 21, got nil")
    }
}

func TestFactorial7(t *testing.T) {
    // Large input far beyond the limit
    _, err := fact(100)
    if err == nil {
        t.Errorf("Expected error for large input, got nil")
    }
}

// ===================================== Natural Log Testcases =========================================

func TestNaturalLog1(t *testing.T) {
    // The log of 1 is always 0
    result, err := naturalLog(1)
    if err != nil || result != 0 {
        t.Errorf("Expected 0, got %g (Error: %v)", result, err)
    }
}

func TestNaturalLog2(t *testing.T) {
    // The log of Euler's number (e) is 1
    result, err := naturalLog(math.E)
    if err != nil || math.Abs(result-1) > 1e-9 {
        t.Errorf("Expected 1, got %g", result)
    }
}

func TestNaturalLog3(t *testing.T) {
    // Test for zero (undefined/negative infinity)
    _, err := naturalLog(0)
    if err == nil {
        t.Errorf("Expected error for input 0, got nil")
    }
}

func TestNaturalLog4(t *testing.T) {
    // Test for negative numbers
    _, err := naturalLog(-5.5)
    if err == nil {
        t.Errorf("Expected error for negative input, got nil")
    }
}

func TestNaturalLog5(t *testing.T) {
    // Test for a value between 0 and 1 (should be negative)
    result, err := naturalLog(0.5)
    if err != nil || result >= 0 {
        t.Errorf("Expected negative result for 0.5, got %g", result)
    }
}

func TestNaturalLog6(t *testing.T) {
    // Large numbers
    result, err := naturalLog(1e10)
    expected := 23.0258509299
    if err != nil || math.Abs(result-expected) > 1e-8 {
        t.Errorf("Expected %g, got %g", expected, result)
    }
}

func TestNaturalLog7(t *testing.T) {
    // Infinity
    result, err := naturalLog(math.Inf(1))
    if err != nil || !math.IsInf(result, 1) {
        t.Errorf("Expected +Inf, got %g", result)
    }
}

func TestNaturalLog8(t *testing.T) {
    // Very small positive number (near the asymptote)
    result, err := naturalLog(1e-300)
    if err != nil || result >= 0 {
        t.Errorf("Expected large negative number, got %g", result)
    }
}

// ================================== Power Testcase =============================================

func TestPow1(t *testing.T) {
    // Standard positive integer power
    result := pow(2, 3)
    if result != 8 {
        t.Errorf("Expected 8, got %g", result)
    }
}

func TestPow2(t *testing.T) {
    // Zero power (Identity): x^0 is always 1
    result := pow(5, 0)
    if result != 1 {
        t.Errorf("Expected 1, got %g", result)
    }
}

func TestPow3(t *testing.T) {
    // Negative exponent (Reciprocal): 2^-2 = 1/(2^2) = 0.25
    result := pow(2, -2)
    if result != 0.25 {
        t.Errorf("Expected 0.25, got %g", result)
    }
}

func TestPow4(t *testing.T) {
    // Fractional exponent (Square Root): 9^0.5 = 3
    result := pow(9, 0.5)
    if result != 3 {
        t.Errorf("Expected 3, got %g", result)
    }
}

func TestPow5(t *testing.T) {
    // Negative base with even exponent (Positive result)
    result := pow(-2, 2)
    if result != 4 {
        t.Errorf("Expected 4, got %g", result)
    }
}

func TestPow6(t *testing.T) {
    // Negative base with odd exponent (Negative result)
    result := pow(-2, 3)
    if result != -8 {
        t.Errorf("Expected -8, got %g", result)
    }
}

func TestPow7(t *testing.T) {
    // Negative base with fractional exponent (Imaginary/NaN)
    // math.Pow returns NaN for negative bases with non-integer exponents
    result := pow(-4, 0.5)
    if !math.IsNaN(result) {
        t.Errorf("Expected NaN for sqrt of negative, got %g", result)
    }
}

func TestPow8(t *testing.T) {
    // Overflow: Large base and large exponent
    result := pow(10, 400)
    if !math.IsInf(result, 1) {
        t.Errorf("Expected +Inf, got %g", result)
    }
}

func TestPow9(t *testing.T) {
    // Zero base with zero exponent (0^0 is generally 1 in programming)
    result := pow(0, 0)
    if result != 1 {
        t.Errorf("Expected 1, got %g", result)
    }
}

// ========================== Some Edge cases ==========================================================
func TestDivisionThreshold(t *testing.T) {
    // Even if validation limits input to 0-9, b could be 0.00000000001
    _, err := divide(5, 0.0000000000001)
    if err == nil {
        t.Errorf("Expected threshold error for near-zero division, got nil")
    }
}

func TestLogOfZero(t *testing.T) {
    // 0 is valid for validateRangeFloat, but invalid for naturalLog
    _, err := naturalLog(0)
    if err == nil {
        t.Errorf("Expected error for Log(0), got nil")
    }
}

func TestValidationBoundary(t *testing.T) {
    // Test the exact upper limit
    if err := validateRangeFloat(9.0); err != nil {
        t.Errorf("Validation failed for 9.0, but it should be valid")
    }
    // Test just outside the upper limit
    if err := validateRangeFloat(9.1); err == nil {
        t.Errorf("Validation passed for 9.1, but it should fail")
    }
    // Test negative input (even if math allows it, your validator shouldn't)
    if err := validateRangeFloat(-1.0); err == nil {
        t.Errorf("Validation passed for -1.0, but it should fail")
    }
}