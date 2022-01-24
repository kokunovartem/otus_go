package hw02unpackstring

func Validate(packedData []Character) error {
	if packedData[0].isNumber || isDoubleNumberInString(packedData) {
		return ErrInvalidString
	}

	return nil
}

func isDoubleNumberInString(packedData []Character) bool {
	for i, character := range packedData {
		if i < len(packedData)-1 && !character.isShielded && character.isNumber && packedData[i+1].isNumber {
			return true
		}
	}
	return false
}
