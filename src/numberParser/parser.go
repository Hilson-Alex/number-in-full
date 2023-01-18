// Author: [Hilson-Alex](https://github.com/Hilson-Alex).

// Package numberParser is a package to get the full written name of a
// number in brazilian portuguese.
package numberParser

import (
	"fmt"
	"strconv"
	"strings"
)

// GetNumberInFull takes a positive integer and gets its written
// name in brazilian portuguese
func GetNumberInFull(number uint64) string {
	var groups = groupHundreds(number)
	var parsedNumber = ""
	for index, hundred := range groups {
		var translated = translateGroup(hundred)
		if len(translated) == 0 {
			continue
		}
		var suffix = groupNames[len(groups)-index-1]
		if translated != UNITS.Translate("1") {
			suffix = plural(suffix)
		}
		parsedNumber += " " + translated + " " + suffix
	}
	if len(parsedNumber) == 0 {
		return ZERO
	}
	return normalize(strings.TrimSpace(parsedNumber))
}

// translateGroup gets the name of each group of hundreds
func translateGroup(group string) string {
	var digits = fmt.Sprintf("%03s", group)
	var translated = make([]string, 0, 3)
	translated = appendNonEmpty(translated, HUNDREDS.Translate(string(digits[0])))
	if digits[1] == '1' {
		translated = appendNonEmpty(translated, TEN_DIGITS.Translate(digits[1:]))
	} else {
		translated = appendNonEmpty(translated, DOZENS.Translate(string(digits[1])))
		translated = appendNonEmpty(translated, UNITS.Translate(string(digits[2])))
	}
	return strings.Join(translated, CONJUNCTION)
}

// groupHundreds divide a number into periods of hundreds.
// for example: with 185900 as an input, it should return
// ["185", "900"]
func groupHundreds(number uint64) []string {
	var groups = make([]string, 0, 7)
	var numberText = strconv.FormatUint(number, 10)
	for numberText != "" {
		var splitIndex = len(numberText) - 3
		if splitIndex < 0 {
			splitIndex = 0
		}
		groups = prepend(groups, numberText[splitIndex:])
		numberText = numberText[:splitIndex]
	}
	return groups
}

// prepend adds an element to the beginning of a slice
func prepend[T any](slice []T, element T) (newSlice []T) {
	newSlice = append(slice[:1], slice...)
	newSlice[0] = element
	return
}

// appendNonEmpty append non-empty strings to the slice and return.
func appendNonEmpty(slice []string, value string) []string {
	if len(value) == 0 {
		return slice
	}
	return append(slice, value)
}
