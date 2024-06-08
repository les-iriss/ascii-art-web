package fs

func AsciiArtFs(text string, banner string) string {
	font := SplitFile("banners/" + banner + ".txt")
	words := SplitText(text)
	return Writer(words, font)
}
