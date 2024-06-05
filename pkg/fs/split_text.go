package fs

import (
	"strings"
)

func SplitText(text string) []string {
	var words []string
	new_text := text
	for strings.Contains(new_text, "\\n") {
		index := strings.Index(new_text,"\\n")
		if index < 0 {
			break
		}
		words = append(words,new_text[:index])
		new_text = new_text[index+2:]
	}
	words = append(words,new_text)
	return words
}
