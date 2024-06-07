package fs

import (
	"fmt"
	"os"
	"strings"
)

func SplitFile(file_name string) [][]string {
	var result [][]string
	// read file content
	content, err_read := os.ReadFile(file_name)
	if err_read != nil {
		// handle error
		fmt.Println("Read file error!")
	}
	// convert content from bytes to string
	txt := string(content)
	txt = strings.ReplaceAll(txt, "\r","")
	// split content file to chars
	slice := strings.Split(txt[1:], "\n\n")
	// split chars to lines
	for _, item := range slice {
		new_item := strings.Split(item, "\n")
		result = append(result, new_item)
	}
	return result
}
