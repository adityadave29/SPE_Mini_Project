package main
//  comment
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFloat(reader *bufio.Reader) (float64, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	value, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, fmt.Errorf("Entered number is invalid.")
	}

	return value, nil
}

func readInt(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid number")
	}

	return value, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n\nScientific Calculator (SPE Mini Project - MT2025013)")
		fmt.Println("1. Add           2. Subtract     3. Multiply")
		fmt.Println("4. Divide        5. Square Root  6. Factorial")
		fmt.Println("7. Natural log   8. Power        9. Exit")
		fmt.Print("Enter the choice: ")

		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)

		choice, err := strconv.Atoi(choiceInput)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid number.")
			continue
		}

		switch choice {

		case 1, 2, 3, 4, 8:
			fmt.Print("Enter first number: ")
			number1, err := readFloat(reader)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}

			fmt.Print("Enter second number: ")
			number2, err := readFloat(reader)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}

			switch choice {
			case 1:
				fmt.Printf("Summation of %.2f and %.2f is %.2f\n",
					number1, number2, add(number1, number2))
			case 2:
				fmt.Printf("Subtraction of %.2f and %.2f is %.2f\n",
					number1, number2, subtract(number1, number2))
			case 3:
				fmt.Printf("Multiplication of %.2f and %.2f is %.2f\n",
					number1, number2, multiply(number1, number2))
			case 4:
				result, err := divide(number1, number2)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				fmt.Printf("Division of %.2f by %.2f is %.2f\n",
					number1, number2, result)
			case 8:
				fmt.Printf("Power of %.2f raised to %.2f is %.2f\n",
					number1, number2, pow(number1, number2))
			}

		case 5:
			fmt.Print("Enter a number: ")
			number, err := readFloat(reader)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}

			result, err := squareRoot(number)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("Square root of %.2f is %.2f\n", number, result)

		case 6:
			fmt.Print("Enter a number: ")
			factNumber, err := readInt(reader)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}

			result, err := fact(factNumber)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("Factorial of %d is %d\n", factNumber, result)

		case 7:
			fmt.Print("Enter a number: ")
			number, err := readFloat(reader)
			if err != nil {
				fmt.Println("Invalid number.")
				continue
			}

			result, err := naturalLog(number)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Printf("Natural logarithm of %.2f is %.4f\n", number, result)

		case 9:
			fmt.Println("Exiting Calculator...")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}