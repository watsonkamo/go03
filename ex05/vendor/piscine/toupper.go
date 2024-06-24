package piscine

func ToUpper(s string) string {
	newStr := ""
	for _, r := range s {
		if r >= 'a' && r <= 'z' {
			newStr += string(r - 32)
		} else {
			newStr += string(r)
		}
	}
	return newStr
}
