package piscine

func Index(s string, toFind string) int {
	for i, c := range s {
		if c == rune(toFind[0]) {
			return i
		}
	}
	return -1
}
