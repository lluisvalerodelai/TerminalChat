package TUI

type Box struct {
	Bordered bool

	Children []*Box
	parent   *Box
	childID  int
	x, y     int
}

func (s Box) RenderAll(height, width int) {
	boxRenderer := Renderer{s, height, width}

	if s.Bordered {
		boxRenderer.DrawBorder()
	}

	if len(s.Children) == 0 {
		return
	}

	for _, child := range s.Children {
		child.x = s.x + 1
		child.y = s.y + 1

		child.RenderAll(height-2, width-2)
	}

}

func (s *Box) AddChild(child *Box) {
  child.parent = s
	s.Children = append(s.Children, child)
}
