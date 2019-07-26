package cmd

import "github.com/mgutz/ansi"

var (
	// bold turns string into a ANSI bold string.
	bold = ansi.ColorFunc("white+h")
	// red turns string into a ANSI red string.
	red = ansi.ColorFunc("red+b")
)
