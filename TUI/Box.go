package TUI

type Box struct {
	Bordered bool

	Children []*Box
	parent   *Box
	childID  int
	x, y     int
}

func (s Box) RenderAll(width, height int) {
	boxRenderer := Renderer{s, width - 1, height - 1} //adjust for width/height starting at 0 but terminal starting at 1

	if s.Bordered {
		boxRenderer.DrawBorder()
	}

	if len(s.Children) == 0 {
		return
	}

	for _, child := range s.Children {
		child.x = s.x + 1
		child.y = s.y + 1

		//height and width must be decreased by 2 because x,y increases by 1
		//we lower the box by 1, then bring it up 2 so that overall its 1 smaller
		child.RenderAll(width-2, height-2)
	}

}

func (s *Box) AddChild(child *Box) {
	child.parent = s
	s.Children = append(s.Children, child)
}
