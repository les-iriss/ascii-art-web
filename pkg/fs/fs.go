package fs

func AsciiArtFs(text string, banner string) string {
	font := SplitFile("banners/"+ banner + ".txt")
	words := SplitText(text)
	result := Writer(words, font)
	return result
}
