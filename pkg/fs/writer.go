package fs

import (
	"fmt"
)

func Writer(words []string, font [][]string) string {
	var result string
	var last_line = false
	for _,word := range words {
		if word == "\\n" {
			result += "\n"
			continue
		}
		if word == "" { continue }
		last_line = true
		for i := 0; i < 8; i++ {
			for _,char := range word {
				if char > '~' || char < ' ' {
					fmt.Println("400 bad request!")
					return ""
				}
				result += font[int(char)-32][i]
			}
			if i < 7 { result += "\n" }
		}
	}
	if last_line { result += "\n" }
	return result
}
