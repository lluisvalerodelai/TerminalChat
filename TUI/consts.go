package main

const (

	/*
		If it starts with f, its a foreground color
		If it starts with b, its a background color
	*/
	fBlack   = 30
	bBlack   = 40
	fRed     = 31
	bRed     = 41
	fGreen   = 32
	bGreen   = 42
	fYellow  = 33
	bYellow  = 43
	fBlue    = 34
	bBlue    = 44
	fMagenta = 35
	bMagenta = 45
	fCyan    = 36
	bCyan    = 46
	fWhite   = 37
	bWhite   = 47
	fDefault = 39
	bDefault = 49
	Reset   = 0 //reset resets the color to terminal default
)

const (
  aCursorHome = "[H"
  aResetDefaults = "[0m"
  aClearScreen = "[2J"
)
