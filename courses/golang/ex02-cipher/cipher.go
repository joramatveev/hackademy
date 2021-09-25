package cipher

type Cipher interface {
	Encode(string) string
	Decode(string) string
}

type Shift struct {
	num int
}

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(n int) Cipher {
	if (n >= 1 && n <= 25) || (n <= -1 && n >= -25) {
		var cipher Cipher = &Shift{n}
		return cipher
	} else {
		return nil
	}
}

func isAlpha(chr rune) int {
	if chr >= 'a' && chr <= 'z' {
		return 1
	}
	if chr >= 'A' && chr <= 'Z' {
		return 2
	}
	return 0
}

func toLower(chr rune) rune {
	return chr - 'A' + 'a'
}

func getPrev(chr rune, n int) string {
	result := int(chr) - n
	if result < 'a' {
		result = result + 26
	} else if result > 'z' {
		result = result - 26
	}
	return string(rune(result))
}

func getNext(chr rune, n int) string {
	result := int(chr) + n
	if result > int('z') {
		result = result - 26
	} else if result < int('a') {
		result = result + 26
	}
	return string(rune(result))
}

func (cipher Shift) Encode(input string) string {
	n := cipher.num
	result := ""
	for _, v := range input {
		valueOfAlpha := isAlpha(v)
		if valueOfAlpha > 0 {
			if valueOfAlpha == 2 {
				result += getNext(toLower(v), n)
			} else {
				result += getNext(v, n)
			}
		}
	}
	return result
}

func (cipher Shift) Decode(input string) string {
	n := cipher.num
	result := ""
	for _, v := range input {
		result += getPrev(v, n)
	}

	return result
}
