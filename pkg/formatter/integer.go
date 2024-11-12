package formatter

import (
	"fmt"
	"strings"
)

func FormatInteger(num int) string {
	s := fmt.Sprintf("%d", num)
	var result strings.Builder
	length := len(s)

	for i, digit := range s {
		if i > 0 && (length-i)%3 == 0 {
			result.WriteRune(',')
		}
		result.WriteRune(digit)
	}

	return result.String()
}
