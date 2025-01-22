package TUI

import (
	"log"

	term "golang.org/x/term"
)

func prepTerm() {
	doAnsii(aClearScreen)   //clear screen
	doAnsii(aCursorHome)    //move cursor to 0,0
	doAnsii(aResetDefaults) //set all values to default
	doAnsii(aCursorInvisible)
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

func resetTerm() {
	doAnsii("[" + "0" + "m") //reset defaults
	doAnsii(aClearScreen)    //clear (need to reset before, else we clear with the color)
	doAnsii(aCursorVisible)  //clear (need to reset before, else we clear with the color)
}
