package TUI

import "fmt"

type Renderer struct {
  Box

	width, height int //Width and height are relative to the Box container
}


func (r Renderer) Render() {

	if r.Bordered {
		r.DrawBorder()
	}

  for child := range r.Children {
    fmt.Printf("Child: +%v", child)
  }

}

func (r Renderer) DrawBorder() {
	r.CursorHome()

	r.put('┌', r.y, r.x)
	r.put('┐', r.y, r.x + r.width)

	r.put('└', r.y + r.height, r.x)
	r.put('┘', r.y + r.height, r.x + r.width)

	// r.hLine(0, 2, r.width-1)
	// r.hLine(r.height, 2, r.width-1)
	//
	// r.vLine(0, 2, r.height-1)
	// r.vLine(r.width, 2, r.height-1)

}

func (r Renderer) vLine(col, row1, row2 int) {

	if row1 > row2 {
		row1, row2 = row2, row1 //swaperoo
	}

	for i := row1; i <= row2; i++ {
		r.put('│', i, col)
	}

}

func (r Renderer) hLine(row int, col1 int, col2 int) {

	if col1 > col2 {
		col1, col2 = col2, col1 //swaperoo
	}

	for i := col1; i <= col2; i++ {
		r.put('─', row, i)
	}

}

func (r Renderer) put(c rune, row, col int) {
	r.to(row, col)
	fmt.Printf("%c\r", c) //print the character, then carraige return move cursor back one col
}

func (r Renderer) to(row, col int) {
	action := fmt.Sprintf("[%v;%vH", row, col)
	doAnsii(action)
}

func (r Renderer) CursorHome() {
	//TODO: make cursorHome go to the home position relative to the box, not global
  //its relative 0 is going to be the parents 0, plus its height offset
}
