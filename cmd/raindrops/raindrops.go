package raindrops

import "strconv"

func Convert(number int) string {
	/**
	  has 3 as a factor, add 'Pling' to the result.
	  has 5 as a factor, add 'Plang' to the result.
	  has 7 as a factor, add 'Plong' to the result.
	  does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.
	*/
	s := ""
	if number%3 == 0 {
		s += "Pling"
	}
	if number%5 == 0 {
		s += "Plang"
	}
	if number%7 == 0 {
		s += "Plong"
	}
	if s == "" {
		s = strconv.Itoa(number)
	}
	return s
}
