package factorial

import (
    "errors"
)

func oldFactorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * oldFactorial(n-1)
	}
}

func Factorial(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("factorial of negative number is undefined")
    }
    if n == 0 {
        return 1, nil
    }
    result := 1
    for i := n; i > 0; i-- {
        result *= i
    }
    return result, nil
}
