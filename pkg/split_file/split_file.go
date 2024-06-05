package split_file

import (
	"os"
	"fmt"
)

func SplitFile(file_name string) []string {
	var result []string
	file, err_open := os.Open(file_name)
	if err_open != nil {
		fmt.Println("Open file error!")
		return []string{"error"}
	}
	content,err_read := os.ReadFile(file_name)
	if err_read != nil {
		fmt.Println("Read file error!")
	}
	defer file.Close()
	fmt.Println(string(content))
	return result
}
