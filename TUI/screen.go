package TUI

import (
	"fmt"
)

type Screen struct {
	width, height int
	color         string //stores either foreground or background
}

func (s Screen) horizontalLine(row int, col1 int, col2 int) {

	switch {
	case col1 < 0:
		col1 = 0
	case col2 > s.width:
		col2 = s.width
	}

  for i := col1; i <= col2; i++ {
    s.put('─', row, i)
  }
}

func (s Screen) resetScreen() {
	doAnsii("[" + "0" + "m") //reset defaults
	doAnsii(aClearScreen)    //clear (need to reset before, else we clear with the color)
	doAnsii(aCursorVisible)    //clear (need to reset before, else we clear with the color)
}

func (s Screen) createOutline(outlineColor string, outlineVal rune) {
	//go to top corner (0,0)
	doAnsii(aCursorHome)

	//set color to foreground color
	s.color = outlineColor
	doAnsii("[" + s.color + "m") //set color to this

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

// draw a fine-line box outline
// columns translate to x, rows translate to y
func (s Screen) createRect(boxColor string, row1, col1, row2, col2 int) {

	switch {
	case row1 > row2:
		row1, row2 = row2, row1 //swaperoo
	case col1 > col2:
		col1, col2 = col2, col1
	}

	switch {
	case row1 < 0:
		row1 = 0
	case row2 > s.height:
		row2 = s.height
	case col1 < 0:
		col1 = 0
	case col2 > s.width:
		col2 = s.width
	}

	//go to top corner (0,0)
	doAnsii(aCursorHome)

	s.color = boxColor
	doAnsii("[" + s.color + "m") //set color to this

	//draw corners

  //we dont go until row2 - 1 because we already manually draw that corner
	for i := row1; i < row2; i++ {
		s.put('│', i, col1)
		s.put('│', i, col2)
	}

  //same here
	for i := col1; i < col2; i++ {
		s.put('─', col1, i)
		s.put('─', col2, i)
	}

	s.put('┌', row1, col1)
	s.put('└', row2, col1)
	s.put('┐', row1, col2)
	s.put('┘', row2, col2)

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
