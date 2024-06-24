package piscine

func ToLower(s string) string {
	newStr := ""
	for _, r := range s {
		if r >= 'A' && r <= 'Z' {
			newStr += string(r + 32)
		} else {
			newStr += string(r)
		}
	}
	return newStr
}
