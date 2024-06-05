package fs

import (
	"fmt"
)

func Writer(words []string, font [][]string) string {
	var result string
	for _,word := range words {
		if word == "\\n" {
			result += "\n"
			continue
		}
		if word == "" { continue }
		for i := 0; i < 8; i++ {
			for _,char := range word {
				if char > '~' || char < ' ' {
					fmt.Println("400 bad request!")
					return ""
				}
				result += font[int(char)-32][i]
			}
			result += "\n"
		}
	}
	return result
}
