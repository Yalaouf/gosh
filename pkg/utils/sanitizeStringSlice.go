package utils

import (
	"slices"
	"strings"
)

func SanitizeInput(input string) []string {
	return slices.DeleteFunc(strings.Split(input, " "), func(element string) bool {
		return element == ""
	})
}
