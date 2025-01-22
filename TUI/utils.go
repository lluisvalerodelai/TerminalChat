package TUI

import (
	"fmt"
)

type CHILD_ID int

// wrapper for fmt.Print()
func doAnsii(seq string) {
	fmt.Print("\033" + seq)
}


