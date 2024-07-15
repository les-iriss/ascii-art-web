package fs

// AsciiArtFs generates ASCII art for the given text using the specified banner font.
func AsciiArtFs(text string, banner string) string {
	font := SplitFile("banners/" + banner + ".txt")
	words := SplitText(text)
	return Writer(words, font)
}
