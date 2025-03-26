package iteration

import "strings"

const repCount = 5

func Repeat(char string, n int) string {
	var repeated strings.Builder
	for i := 0; i < repCount; i++ {
		repeated.WriteString(char)
	}
	return repeated.String()
}
