package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, n int) (float64, error) {
	amount, err := fc.FodderAmount(n)
	if err != nil {
		return 0, err
	}
	factor, err := fc.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return amount * factor / float64(n), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
	if n <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fc, n)
}

// TODO: define the 'ValidateNumberOfCows' function
// => "-5 cows are invalid: there are no negative cows"
// {number of cows} cows are invalid: {custom message}
type InvalidCowsError struct {
	num int
}

func (receiver *InvalidCowsError) Error() string {
	if receiver.num == 0 {
		return fmt.Sprintf("%d cows are invalid: %s", receiver.num, "no cows don't need food")
	}
	return fmt.Sprintf("%d cows are invalid: %s", receiver.num, "there are no negative cows")
}

func ValidateNumberOfCows(n int) error {
	if n > 0 {
		return nil
	}
	return &InvalidCowsError{n}
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
