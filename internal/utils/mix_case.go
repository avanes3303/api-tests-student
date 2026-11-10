package utils

import "unicode"

// MixCase чередует нижний и верхний регистр букв, включая кириллицу.
func MixCase(text string) string {
	letters := []rune(text)
	upper := false

	for i, letter := range letters {
		if !unicode.IsLetter(letter) {
			continue
		}

		if upper {
			letters[i] = unicode.ToUpper(letter)
		} else {
			letters[i] = unicode.ToLower(letter)
		}
		upper = !upper
	}

	return string(letters)
}
