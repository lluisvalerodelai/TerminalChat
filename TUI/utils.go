package TUI

import (
	"fmt"
	"log"

	term "golang.org/x/term"
)

// wrapper for fmt.Print()
func doAnsii(seq string) {
	fmt.Print("\033" + seq)
}

func createTerm(width, height int) Screen {
	doAnsii(aClearScreen)   //clear screen
	doAnsii(aCursorHome)    //move cursor to 0,0
	doAnsii(aResetDefaults) //set all values to default
  doAnsii(aCursorInvisible)

	// cursor := Cursor{0, 0}
	terminal := Screen{width, height, "0"}

	return terminal
}

func getTermDimensions(termFD int) (width, height int) {
	width, height, err := term.GetSize(termFD)

	if err != nil {
		log.Fatal(err)
	}

	return width, height
}

// setTermRaw sets the terminal to raw state and handles the error
// it will return the oldState pointer
func setTermRaw(fd int) *term.State {
	oldState, err := term.MakeRaw(fd)

	if err != nil {
		log.Fatal(err)
	}

	return oldState
}
