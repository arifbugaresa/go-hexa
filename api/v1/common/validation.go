package common

func IsStringEmpty(data string) bool {
	if data == "" {
		return true
	}

	return false
}

func IsNumericZeroOrMinus(data int) bool {
	if data < 1 {
		return true
	}

	return false
}

func IsFloatZeroOrMinus(data float64) bool {
	if data < 1 {
		return true
	}

	return false
}