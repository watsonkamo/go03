package piscine

func LastRune(s string) rune {
	r := []rune(s)
	l := 0
	for range s {
		l++
	}
	return r[l-1]
}
