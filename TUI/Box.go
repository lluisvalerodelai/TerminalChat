package TUI

type Box struct {
	Bordered bool

	Children []*Box
	parent   *Box
	childID  CHILD_ID
  x, y     int
}

func (s Box) RenderAll(height, width int) {
	boxRenderer := Renderer{s, height, width}

	if s.Bordered {
		boxRenderer.DrawBorder()
	}

	//allocate height amongst the children
	numChildren := len(s.Children)
	if numChildren == 0 { //base case
		return
	}

	childrenHeight := height / numChildren

	//render the children
	for _, child := range s.Children {
    child.x = s.x + childrenHeight * int(child.childID)
    child.y = s.y + 1

		child.RenderAll(childrenHeight, width-1)
	}

}

func (s Box) AddChild(child *Box) {
	child.parent = &s

	//set the child ID
	if len(s.Children) == 0 {
		child.childID = 0
	} else {
		//child ID increment by 1
		child.childID = s.Children[len(s.Children)-1].childID + 1
	}

	s.Children = append(s.Children, child)
}
