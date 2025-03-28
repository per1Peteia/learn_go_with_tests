package main

import "strings"

func ConvertToRoman(number int) string {
	var result strings.Builder
	for number > 0 {
		switch {
		case number > 9:
			result.WriteString("X")
			number -= 10
		case number > 8:
			result.WriteString("IX")
			number -= 9
		case number > 4:
			result.WriteString("V")
			number -= 5
		case number > 3:
			result.WriteString("IV")
			number -= 4
		default:
			result.WriteString("I")
			number--
		}
	}

	return result.String()
}

