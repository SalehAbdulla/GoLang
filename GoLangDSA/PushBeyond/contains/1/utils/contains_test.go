package utils

import (
	"fmt"
	"testing"
)

func TestSimpleIncrement(t *testing.T) {
	code := "text text text text"
	toFind := "text"
	doesExist := contains(code, toFind)
	fmt.Println(doesExist)
}
