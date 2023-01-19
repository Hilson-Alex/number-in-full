// Author: [Hilson-Alex](https://github.com/Hilson-Alex).

// Package numberParser is a package to get the full written name of a
// number in brazilian portuguese.
package numberParser

import (
	"regexp"
	"strings"
)

// unitPlace is a wrapper struct to simulate a constant map
// for the name of a digit in a unit place.
type unitPlace struct {
	translate map[string]string
}

// Translate gets the written name of the number (in portuguese).
func (place *unitPlace) Translate(number string) string {
	return place.translate[number]
}

// UNITS maps the name of the digits on the units place.
var UNITS = &unitPlace{
	translate: map[string]string{
		"0": "",
		"1": "um",
		"2": "dois",
		"3": "três",
		"4": "quatro",
		"5": "cinco",
		"6": "seis",
		"7": "sete",
		"8": "oito",
		"9": "nove",
	},
}

// TEN_DIGITS maps the name of the digits between 10 and 19,
// as they are named different.
var TEN_DIGITS = &unitPlace{
	translate: map[string]string{
		"10": "dez",
		"11": "onze",
		"12": "doze",
		"13": "treze",
		"14": "quatorze",
		"15": "quinze",
		"16": "dezesseis",
		"17": "dezessete",
		"18": "dezoito",
		"19": "dezenove",
	},
}

// DOZENS maps the name of the digits on the tens place.
var DOZENS = &unitPlace{
	translate: map[string]string{
		"0": "",
		"2": "vinte",
		"3": "trinta",
		"4": "quarenta",
		"5": "cinquenta",
		"6": "sessenta",
		"7": "setenta",
		"8": "oitenta",
		"9": "noventa",
	},
}

// HUNDREDS maps the name of the digits on the hundreds place.
var HUNDREDS = &unitPlace{
	translate: map[string]string{
		"0": "",
		"1": "cem",
		"2": "duzentos",
		"3": "trezentos",
		"4": "quatrocentos",
		"5": "quinhentos",
		"6": "seiscentos",
		"7": "setecentos",
		"8": "oitocentos",
		"9": "novecentos",
	},
}

// ZERO is just the name for zero.
const ZERO = "zero"

// CONJUNCTION is the string to join each digit name.
// It's like "and" in portuguese.
const CONJUNCTION = " e "

// groupNames maps the name for each period. Thousands, millions,
// billions and so on...
var groupNames = [...]string{"", "mil", "milhão", "bilhão", "trilhão", "quatrilhão", "quintilhão"}

// changes a period to its plural form.
func plural(groupName string) string {
	return strings.ReplaceAll(groupName, "ão", "ões")
}

// normalize makes small corrections for the final generated string.
func normalize(parsedNumber string) string {
	parsedNumber = strings.ReplaceAll(parsedNumber, "cem"+CONJUNCTION, "cento"+CONJUNCTION)
	parsedNumber = regexp.MustCompile("(([^e] )|^)um mil").ReplaceAllStringFunc(parsedNumber, func(s string) string {
		return strings.Replace(s, "um ", "", 1)
	})
	var hasHundreds = false
	for _, value := range HUNDREDS.translate {
		if value == "" {
			continue
		}
		if regexp.MustCompile("mil " + value + "$").MatchString(parsedNumber) {
			return strings.Replace(parsedNumber, "mil ", "mil"+CONJUNCTION, 1)
		}
		if regexp.MustCompile("mil " + value).MatchString(parsedNumber) {
			hasHundreds = true
			break
		}
	}
	if !hasHundreds && regexp.MustCompile("mil \\w+").MatchString(parsedNumber) {
		return strings.Replace(parsedNumber, "mil ", "mil"+CONJUNCTION, 1)
	}
	return parsedNumber
}
