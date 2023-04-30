package thefarm

import "fmt"

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	n int
}

func (sne *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sne.n)
}

type DivisionByZeroError struct{}

func (dbz *DivisionByZeroError) Error() string {
	return "division by zero"
}

type NegativeFooderError struct{}

func (nf NegativeFooderError) Error() string {
	return "negative fodder"
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	amount, err := weightFodder.FodderAmount()

	switch {
	case isAmoutPositiveAndCowsNegative(amount, cows):
		return 0, &SillyNephewError{n: cows}
	case isCowsZero(cows):
		return 0, &DivisionByZeroError{}
	case isErrScaleMalfunctionAndAmountPositive(err, amount):
		return amount * 2 / float64(cows), nil
	case isAmountNegativeAndErrorErrScaleMalfunctionOrNil(amount, err):
		return 0, &NegativeFooderError{}
	case isErrorNil(err):
		return 0, err
	default:
		return amount / float64(cows), nil
	}
}

func isAmoutPositiveAndCowsNegative(amount float64, cows int) bool {
	return amount >= 0 && cows < 0
}

func isCowsZero(cows int) bool {
	return cows == 0
}

func isErrScaleMalfunctionAndAmountPositive(err error, amount float64) bool {
	return err == ErrScaleMalfunction && amount > 0
}

func isAmountNegativeAndErrorErrScaleMalfunctionOrNil(amount float64, err error) bool {
	return amount < 0 && (err == ErrScaleMalfunction || err == nil)
}

func isErrorNil(err error) bool {
	return err != nil
}
