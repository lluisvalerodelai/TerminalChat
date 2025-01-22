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

	prepTerm()

	rootBox := &Box{Bordered: true}

	child1 := &Box{Bordered: true}
	child2 := &Box{Bordered: true}
	child3 := &Box{Bordered: true}
	child4 := &Box{Bordered: true}

	// Manually link the hierarchy
	rootBox.AddChild(child1)
	child1.AddChild(child2)
	child2.AddChild(child3)
	child3.AddChild(child4)

	width, height := getTermDimensions(termFD)
	rootBox.x = 1
	rootBox.y = 1

	rootBox.RenderAll(width, height)

	actionBuf := make([]byte, 1)
	for {

		os.Stdin.Read(actionBuf)

		if string(actionBuf[0]) == "q" {
			resetTerm()
			return
		}

	}
}
