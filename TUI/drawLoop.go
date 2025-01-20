package TUI

import (
	"fmt"
	"log"
	"os"
	"strconv"

	term "golang.org/x/term"
)

type Screen struct {
	width, height int
	color         int //stores either foreground or background

	//state (2d array representation of terminal)
}

//because (whatever whatsapp calls it) 2 way secure communication is bloated
//mock whatsapps incredibly fast message speeds and capabilities and stuff

func (s Screen) createOutline(outlineColor int, outlineVal rune) {
	//draw top and bottom row, then loop thorugh 1 -> col - 1

	//go to top corner (0,0)
	doAnsii(aCursorHome)

	//set color to foreground color
	s.color = outlineColor
	doAnsii("[" + strconv.Itoa(s.color) + "m") //set color to this

	//draw top of outline
	for i := 0; i < s.width; i++ {
		s.put(outlineVal, 0, i)        //put it in row 0, col i
		s.put(outlineVal, s.height, i) //put it in row 0, col i
	}

	for i := 1; i < s.height-1; i++ {
		s.put(outlineVal, i, 0)       //put it in row i, col 0
		s.put(outlineVal, i, s.width) //put it in row i, leftmost column
	}

}

// place a character at x,y
func (s Screen) put(c rune, row, col int) {
	s.to(row, col)
	fmt.Printf("%c\r", c) //print the character, then carraige return move cursor back one col
}

// move the cursor to x, y
func (s Screen) to(row, col int) {
	action := fmt.Sprintf("[%v;%vH", row, col)
	doAnsii(action)
}

func MainDrawLoop() {

	//setup
	termFD := int(os.Stdin.Fd())

	oldState := setTermRaw(termFD) //old state type is *term.State
	defer term.Restore(termFD, oldState)

	width, height := getTermDimensions(termFD)
	terminal := createTerm(width, height)

	terminal.createOutline(bRed, ' ')

	actionBuf := make([]byte, 1)
	for {

		os.Stdin.Read(actionBuf)
		if string(actionBuf[0]) == "q" {
			return
		}
	}

	//main loop
	/*
	   poll for updates
	   draw updates
	*/
}

func createTerm(width, height int) Screen {
	doAnsii(aClearScreen)   //clear screen
	doAnsii(aCursorHome)    //move cursor to 0,0
	doAnsii(aResetDefaults) //set all values to default

	// cursor := Cursor{0, 0}
	terminal := Screen{width, height, 0}

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
