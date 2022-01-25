package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

type Character struct {
	char       string
	isNumber   bool
	value      int
	isShielded bool
	isAsterisk bool
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(packedString string) (string, error) {
	if len(packedString) == 0 {
		return "", nil
	}

	packedData := toStruct(strings.Split(packedString, ""))
	shieldWithAsterisk(packedData)

	if nil != Validate(packedData) {
		return "", ErrInvalidString
	}

	unpackedData := make([]Character, len(packedString))

	for i, character := range packedData {
		if character.isShielded {
			unpackedData[i] = character
			continue
		}

		if character.isNumber {
			if character.value == 0 {
				unpackedData[i-1].char = ""
			}

			unpackedData[i-1].char = strings.Repeat(packedData[i-1].char, character.value)
		}
		unpackedData[i] = character
	}

	return toString(unpackedData), nil
}

func toStruct(packedSlice []string) []Character {
	packedData := make([]Character, len(packedSlice))
	for i, char := range packedSlice {
		var isNumber bool
		value, err := strconv.Atoi(char)
		if err == nil {
			isNumber = true
		}
		packedData[i] = Character{
			char,
			isNumber,
			value,
			false,
			false,
		}
	}

	return packedData
}

func shieldWithAsterisk(packedData []Character) {
	for i, character := range packedData {
		if character.char == `\` && !character.isShielded {
			packedData[i].isAsterisk = true
			if i < len(packedData)-1 {
				packedData[i+1].isShielded = true
			}
		}
	}
}

func toString(unpackedData []Character) string {
	var result string
	for _, character := range unpackedData {
		if (character.isNumber && !character.isShielded) || (character.isAsterisk && !character.isShielded) {
			continue
		}

		result += character.char
	}

	return result
}
