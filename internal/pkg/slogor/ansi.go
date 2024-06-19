package slogor

// ANSI codes for text styling and formatting.
var (
	// https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_(Select_Graphic_Rendition)_parameters
	reset           = "\033[0m"
	bold            = "\033[1m"
	faint           = "\033[2m"
	underline       = "\033[4m"
	normalIntensity = "\033[22m"
	// https://en.wikipedia.org/wiki/ANSI_escape_code#3-bit_and_4-bit
	fgRed     = "\033[31m"
	fgGreen   = "\033[32m"
	fgYellow  = "\033[33m"
	fgBlue    = "\033[34m"
	fgMagenta = "\033[35m"
	fgCyan    = "\033[36m"
)
