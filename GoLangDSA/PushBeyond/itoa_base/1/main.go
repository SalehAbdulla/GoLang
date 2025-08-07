package main

import "fmt"

const digits = "0123456789ABCDEF"

func ItoaBase(value, base int) string {
	if value == 0 {
		return "0"
	}
	isNegative := value < 0
	var uValue uint64

	if isNegative {
		uValue = uint64(-int64(value))
	} else {
		uValue = uint64(value)
	}
	result := ""
	for uValue > 0 {
		reminder := uValue % uint64(base)
		result = string(digits[reminder]) + result
		uValue /= uint64(base)
	}

	if isNegative {
		return "-" + result
	} else {
		return result
	}

}

func main() {
	fmt.Println(ItoaBase(2, 16))
}