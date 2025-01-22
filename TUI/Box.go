package TUI

type Box struct {
	Bordered bool

	Children []*Box
	parent   *Box
	childID  int
	x, y     int
}

func (s Box) RenderAll(width, height int, ul, ur, dl, dr rune) {

	//grid starts at 1, width,height get counted as if it started at 0
	width = width - 1
	height = height - 1

	boxRenderer := Renderer{s, width, height} //adjust for width/height starting at 0 but terminal starting at 1

	if s.Bordered {
		if s.childID != 0 {
			boxRenderer.DrawBorder('┬', ur, '┴', dr)
		} else {
			boxRenderer.DrawBorder(ul, ur, dl, dr)
		}

	}

	var offset int = 1

	if len(s.Children) == 0 { //base case
		return
	} else {
		offset = width / len(s.Children)
	}

	//here just responsible for drawing the children
	for n, child := range s.Children {

		child.x = s.x + 1 + (offset * n)
		child.y = s.y + 1

		child.RenderAll((width/(n+1))-1, height-1, ul, ur, dl, dr)
	}

}

func (s *Box) AddChild(child *Box) {
	child.parent = s

	if len(s.Children) == 0 {
		child.childID = 0
	} else {
		//child ID increment by 1
		child.childID = s.Children[len(s.Children)-1].childID + 1
	}

	s.Children = append(s.Children, child)
}
