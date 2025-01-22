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

	//draw lines before corners so we dont have to do off by 1 adjustments
	//instead just draw over them

	r.hLine(r.y, r.x, r.x+r.width)
	r.hLine(r.y+r.height, r.x, r.x+r.width)
	r.vLine(r.x, r.y, r.y+r.height)
	r.vLine(r.x+r.width, r.y, r.y+r.height)

	r.put('┌', r.y, r.x)
	r.put('┐', r.y, r.x+r.width)

	r.put('└', r.y+r.height, r.x)
	r.put('┘', r.y+r.height, r.x+r.width)
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
