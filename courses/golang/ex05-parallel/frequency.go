package letter

func Frequency(text string) map[rune]int {
	result := map[rune]int{}

	for _, i := range text {
		result[i]++
	}

	return result
}

func ConcurrentFrequency(text []string) map[rune]int {
	result := map[rune]int{}

	j := make(chan map[rune]int)
	for _, i := range text {
		go calc(i, j)
	}

	for range text {
		for k, i := range <-j {
			result[k] += i
		}
	}

	return result
}

func calc(text string, ch chan map[rune]int) {
	ch <- Frequency(text)
}
