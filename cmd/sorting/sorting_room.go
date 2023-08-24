package sorting

import (
	"fmt"
	"strconv"
)

// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	// Output: This is the number -12.3
	return fmt.Sprintf("This is the number %.1f", f)
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
// Output: This is a box containing the number 12.0
func DescribeNumberBox(nb NumberBox) string {
	return fmt.Sprintf("This is a box containing the number %.1f", float64(nb.Number()))
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	number, ok := fnb.(FancyNumber)
	if ok {
		var n, _ = strconv.Atoi(number.Value())
		return n
	} else {
		return 0
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
// Output: This is a fancy box containing the number 10.0
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	return fmt.Sprintf("This is a fancy box containing the number %.1f", float64(ExtractFancyNumber(fnb)))
}

// DescribeAnything should return a string describing whatever it contains.
/**
This is the main function Jen needs which takes any input (the empty interface means any value at all: interface{}).
DescribeAnything should delegate to the other functions based on the type of the value passed in. More specifically:
    int and float64 should both delegate to DescribeNumber
    NumberBox should delegate to DescribeNumberBox
    FancyNumberBox should delegate to DescribeFancyNumberBox
    anything else should result in "Return to sender"
fmt.Println(DescribeAnything(numberBoxContaining{12.345}))
// Output: This is a box containing the number 12.3
fmt.Println(DescribeAnything("some string"))
// Output: Return to sender


*/
func DescribeAnything(i interface{}) string {
	switch t := i.(type) {
	case int:
		return DescribeNumber(float64(t))
	case float64:
		return DescribeNumber(t)
	case NumberBox:
		return DescribeNumberBox(t)
	case FancyNumberBox:
		return DescribeFancyNumberBox(t)
	default:
		return "Return to sender"
	}
}
