// mystrings pacakge
// this is a simple strings package

package mystrings

import (
	"fmt"
	"strings"
)

// ConcatenateUsingPlusOperator using the + operator (inefficient)
// this takes n int as an argument and returns a string
func ConcatenateUsingPlusOperator(n int) string {
	var result string
	for i := 0; i < n; i++ {
		result += fmt.Sprintf("%d", i)
	}
	return result
}

// ConcatenateUsingBuilder using strings.Builder (efficient)
func ConcatenateUsingBuilder(n int) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(fmt.Sprintf("%d", i))
	}
	return builder.String()
}
