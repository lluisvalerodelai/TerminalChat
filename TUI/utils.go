package TUI

import "fmt"

//wrapper for fmt.Print()
func doAnsii(seq string) {
  fmt.Print("\033" + seq)
}
