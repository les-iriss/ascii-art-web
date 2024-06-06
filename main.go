package main

import (
	"fmt"
	fs "ascii-art-web/pkg/fs"
)

func main() {
	var text, banner = "\\nhello\\n\\nman\\n", "standard"
	ascii_fs := fs.AsciiArtFs(text, banner)
	fmt.Print(ascii_fs)
}

