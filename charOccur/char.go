package charOccur

func CharacterOccurence(input string) map[rune]int {
	m := map[rune]int{}
	for _, ch := range input {
		m[ch]++
	}

	return m
}
