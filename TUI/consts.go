package TUI

const (

	/*
		If it starts with f, its a foreground color
		If it starts with b, its a background color
	*/
	fBlack      = "30"
	bBlack      = "40"
	fRed        = "31"
	bRed        = "41"
	fGreen      = "32"
	bGreen      = "42"
	fYellow     = "33"
	bYellow     = "43"
	fBlue       = "34"
	bBlue       = "44"
	fMagenta    = "35"
	bMagenta    = "45"
	fCyan       = "36"
	bCyan       = "46"
	fWhite      = "37"
	bWhite      = "47"
	fDefault    = "39"
	bDefault    = "49"
	termDefault = "0" //use the terminal default (no color)
)

const (
	aCursorHome      = "[H"
	aResetDefaults   = "[0m"
	aClearScreen     = "c"
	aCursorVisible   = "[?25h"
	aCursorInvisible = "[?25l"
)

const Header = 
`
 ______   __                   __      __       __  __                   __               
 /      \ |  \                 |  \    |  \  _  |  \|  \                 |  \              
|  $$$$$$\| $$____    ______  _| $$_   | $$ / \ | $$| $$____    ______  _| $$_     _______ 
| $$   \$$| $$    \  |      \|   $$ \  | $$/  $\| $$| $$    \  |      \|   $$ \   /       \
| $$      | $$$$$$$\  \$$$$$$\\$$$$$$  | $$  $$$\ $$| $$$$$$$\  \$$$$$$\\$$$$$$  |  $$$$$$$
| $$   __ | $$  | $$ /      $$ | $$ __ | $$ $$\$$\$$| $$  | $$ /      $$ | $$ __  \$$    \ 
| $$__/  \| $$  | $$|  $$$$$$$ | $$|  \| $$$$  \$$$$| $$  | $$|  $$$$$$$ | $$|  \ _\$$$$$$\
 \$$    $$| $$  | $$ \$$    $$  \$$  $$| $$$    \$$$| $$  | $$ \$$    $$  \$$  $$|       $$
  \$$$$$$  \$$   \$$  \$$$$$$$   \$$$$  \$$      \$$ \$$   \$$  \$$$$$$$   \$$$$  \$$$$$$$ 
                                                                                           
`  
