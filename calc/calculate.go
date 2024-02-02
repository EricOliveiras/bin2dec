package calc

import (
	"errors"
	"fmt"
	"strconv"
)

// ValidateBinary validates a binary string, allowing up to 8 digits and only 0 and 1.
func validateBinary(binaryNumber string) error {
	if len(binaryNumber) > 8 {
		return errors.New("only up to 8 digits are allowed")
	}

	for _, char := range binaryNumber {
		if string(char) != "1" && string(char) != "0" {
			return errors.New("only 0 and 1 are allowed")
		}
	}

	return nil
}

// ReadInputNumber prompts the user for input and prints the result.
func ReadInputNumber() {
	for {
		var binaryNumber string

		fmt.Println("Enter the binary number:")
		fmt.Scanln(&binaryNumber)

		err := validateBinary(binaryNumber)
		if err != nil {
			fmt.Println("Validation Error:", err)
			return
		}

		decimalNumber, err := strconv.ParseInt(binaryNumber, 2, 0)
		if err != nil {
			fmt.Println("Calculation Error:", err)
			return
		}

		fmt.Printf("The decimal number for the binary number %s is: %d\n\n", binaryNumber, decimalNumber)

		var yesOrNo string
		fmt.Println("Would you like to enter another number? (y / n)")
		fmt.Scan(&yesOrNo)

		if yesOrNo != "y" && yesOrNo != "Y" {
			break
		}
	}
}
