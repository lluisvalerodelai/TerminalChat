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

  rootBox := &Box{}
  rootBox.Bordered = true

  child1 := &Box{}
  child1.Bordered = true

  child2 := &Box{}
  child2.Bordered = true

  rootBox.AddChild(child1)
  rootBox.AddChild(child2)

  height, width := getTermDimensions(termFD)
  rootBox.RenderAll(height, width)

	actionBuf := make([]byte, 1)
	for {

		os.Stdin.Read(actionBuf)

		if string(actionBuf[0]) == "q" {
      resetTerm() 
			return
		}

	}
}
