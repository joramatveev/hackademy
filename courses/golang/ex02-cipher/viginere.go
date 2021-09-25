package cipher

type Viginere struct {
	str string
}

func NewVigenere(str string) Cipher {
	if !hasBlockedCharacters(str) {
		var result Cipher = &Viginere{str}
		return result
	}
	return nil
}

func (cipher Viginere) Encode(input string) string {
	key := cipher.str
	i := 0
	result := ""
	for _, v := range input {
		vAlph := isAlpha(v)
		if vAlph > 0 {
			if i >= len(key) {
				i = 0
			}
			x := int(key[i] - 'a')
			i++
			if vAlph == 2 {
				result += getNext(toLower(v), x)
			} else {
				result += getNext(v, x)
			}
		}
	}
	return result
}

func (cipher Viginere) Decode(input string) string {
	j := 0
	result := ""
	key := cipher.str
	for i := 0; i < len(input); i++ {
		if j >= len(key) {
			j = 0
		}
		delta := key[j] - 'a'
		j++
		newStr := input[i] - delta
		if newStr < 'a' {
			newStr += 26
		}
		result += string(newStr)
	}
	return result
}

func hasBlockedCharacters(str string) bool {
	thereWasNoA := false
	for _, chr := range str {
		if chr < 'a' || chr > 'z' {
			return true
		}
		if chr != 'a' {
			thereWasNoA = true
		}
	}
	return !thereWasNoA
}
