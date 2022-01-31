package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

const asteriskRune = 92

type RuneSymbol struct {
	value      rune
	isShielded bool
}

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	if len(input) == 0 {
		return "", nil
	}

	var (
		output         strings.Builder
		prevRuneSymbol RuneSymbol
	)

	inputData := make([]RuneSymbol, len([]rune(input)))
	for i, r := range []rune(input) {
		inputData[i] = RuneSymbol{r, false}
	}
	inputData = shieldWithAsterisk(inputData)

	for i, runeSymbol := range inputData {
		validateErr := validate(i, prevRuneSymbol, runeSymbol)
		if validateErr != nil {
			return "", validateErr
		}

		if isDigit(runeSymbol) {
			num, _ := strconv.Atoi(string(runeSymbol.value))
			if num == 0 {
				outputString := output.String()
				_, size := utf8.DecodeLastRuneInString(outputString)
				output.Reset()
				output.WriteString(outputString[:len(outputString)-size])
				prevRuneSymbol = runeSymbol
				continue
			}

			output.Write([]byte(strings.Repeat(string(prevRuneSymbol.value), num-1)))
			prevRuneSymbol = runeSymbol
			continue
		}

		prevRuneSymbol = runeSymbol

		if !isDigit(runeSymbol) && !isAsterisk(runeSymbol) {
			output.WriteRune(runeSymbol.value)
		}
	}

	return output.String(), nil
}

func shieldWithAsterisk(inputData []RuneSymbol) []RuneSymbol {
	for i, runeSymbol := range inputData {
		if isAsterisk(runeSymbol) && i < len(inputData)-1 {
			inputData[i+1].isShielded = true
		}
	}

	return inputData
}

func validate(index int, prevRuneSymbol, runeSymbol RuneSymbol) error {
	conditions := []bool{
		index == 0 && isDigit(runeSymbol),
		isDigit(prevRuneSymbol) && isDigit(runeSymbol),
	}

	for _, condition := range conditions {
		if condition {
			return ErrInvalidString
		}
	}

	return nil
}

func isAsterisk(runeSymbol RuneSymbol) bool {
	return runeSymbol.value == asteriskRune && !runeSymbol.isShielded
}

func isDigit(runeSymbol RuneSymbol) bool {
	return unicode.IsDigit(runeSymbol.value) && !runeSymbol.isShielded
}
