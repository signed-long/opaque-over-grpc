package cli

func isAllowedChar(char rune) bool {
	isAlpha := ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z')
	isNum := ('0' <= char && char <= '9')
	isUnderscore := (char == '-')
	return (isAlpha || isNum || isUnderscore)
}
