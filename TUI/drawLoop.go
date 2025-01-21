package TUI

import (
	"os"

	term "golang.org/x/term"
)

func MainDrawLoop() {

	//setup
	termFD := int(os.Stdin.Fd())

	oldState := setTermRaw(termFD) //old state type is *term.State
	defer term.Restore(termFD, oldState)

	width, height := getTermDimensions(termFD)
	terminal := createTerm(width, height)

	terminal.createRect(termDefault, 0, 0, terminal.height, terminal.width)

	terminal.horizontalLine(terminal.height/5, 2, terminal.width-1)
	terminal.putTextRaw(Header, 2, 2)

	actionBuf := make([]byte, 1)
	for {

		os.Stdin.Read(actionBuf)

		if string(actionBuf[0]) == "q" {
			terminal.resetScreen()
			return
		}

	}
}
