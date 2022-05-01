package screen

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

// Screen defines the current layout.
type Screen func(gtx layout.Context) (Screen, layout.Dimensions)

var in = layout.UniformInset(unit.Dp(8)) // Default inset.

func rowInset(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { return in.Layout(gtx, w) }
}

func colInset(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { return in.Layout(gtx, w) }
}
