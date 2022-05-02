package screen

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Screen defines the current layout.
type Screen func(gtx layout.Context) (Screen, layout.Dimensions)

var (
	s      = unit.Dp(5)
	in     = layout.UniformInset(s) // Default inset.
	spacer = layout.Spacer{Width: s, Height: s}
)

func rowInset(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { return in.Layout(gtx, w) }
}
