package main

func DigitRuneToInt(r rune) (int, bool) {
	if r >= '0' && r <= '9' {
		return int(r - '0'), true
	}
	return 0, false
}
