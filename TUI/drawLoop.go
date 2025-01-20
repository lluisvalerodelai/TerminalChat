package main

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

	//draw top of outline
	for i := 0; i < s.width; i++ {
    doAnsii("[" + strconv.Itoa(s.color) + "m") //set color to this
    s.put(outlineVal) 
	}

  for i := 1; i < s.height - 1; i++ {
    s.put(outlineVal) 
  }

  for i:= s.height; i < s.width; i++ {
    s.put(outlineVal)
  }
}

func (s Screen) put(c rune) {
  fmt.Printf("%c\r", c) //print the character, then carraige return move cursor back one col
}


func main() {

	//setup
	termFD := int(os.Stdin.Fd())

	oldState := setTermRaw(termFD) //old state type is *term.State
	defer term.Restore(termFD, oldState)

	width, height := getTermDimensions(termFD)

	terminal := createTerm(width, height)

	terminal.createOutline(fRed, '0')

  for {}

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
