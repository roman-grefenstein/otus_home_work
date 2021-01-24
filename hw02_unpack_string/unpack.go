package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var prevChar rune
	var result strings.Builder
	for i, char := range str {
		if (i == 0 && unicode.IsDigit(char)) || (unicode.IsDigit(char) && unicode.IsDigit(prevChar)) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(char) {
			num, err := strconv.Atoi(string(str[i]))
			if err != nil {
				return "", errors.New("atoi error")
			}
			result.WriteString(strings.Repeat(string(prevChar), num))
		} else {
			if i != 0 && !unicode.IsDigit(prevChar) {
				result.WriteString(string(prevChar))
			}
			if i == len(str)-1 {
				result.WriteString(string(str[i]))
			}
		}

		prevChar = char
	}

	return result.String(), nil
}
