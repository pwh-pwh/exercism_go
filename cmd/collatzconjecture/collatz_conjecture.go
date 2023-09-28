package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0
	if n <= 0 {
		return 0, errors.New("")
	}
	for n != 1 {
		count++
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
	return count, nil
}
