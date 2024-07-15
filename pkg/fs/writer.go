package fs

// Writer generates ASCII art for the given words using the specified font.
func Writer(words []string, font [][]string) string {
	var result string
	last_line := false
	for _, word := range words {
		if word == "\\n" {
			result += "\n"
			continue
		}
		if word == "" {
			continue
		}
		last_line = true
		for i := 0; i < 8; {
			for _, char := range word {
				if char > '~' || char < ' ' {
					continue
				}
				result += font[int(char)-32][i]
			}
			if i < 7 {
				result += "\n"
			}
			i++
		}
	}
	if last_line {
		result += "\n"
	}
	return result
}
