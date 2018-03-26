package shapes

import "github.com/tengrommel/go_live/go_job/design_patterns/Behaviral/strategy/example2"

type TextSquare struct {
	strategy.DrawOutput
}

func (t *TextSquare) Draw() error {
	t.Writer.Write([]byte("Circle"))
	return nil
}
