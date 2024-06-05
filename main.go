package main

import (
	"ascii-art-web/pkg/split_file"
	"fmt"
)

func main() {
	var text, banner = "hello", "standard"
	font := split_file.SplitFile("banners/"+ banner + ".txt")
	fmt.Println(text,font)
}

