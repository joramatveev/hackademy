package downcase

func isUpper(chr string) bool {
	return chr[0] >= 'A' && chr[0] <= 'Z'
}

func Downcase(char string) (string, error) {
	res := ""
	for i := 0; i < len(char); i++ {
		if isUpper(string(char[i])) {
			res += string(char[i] - 'A' + 'a')
		} else {
			res += string(char[i])
		}
	}
	return res, nil
}