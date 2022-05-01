package screen

import (
	"image"

	"gioui.org/layout"
	"gioui.org/unit"
)

// Screen defines the current layout.
type Screen func(gtx layout.Context) (Screen, layout.Dimensions)

var (
	in     = layout.UniformInset(unit.Dp(8)) // Default inset.
	spacer = image.Pt(8, 8)
)

func rowInset(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { return in.Layout(gtx, w) }
}

func spacerLayout(gtx layout.Context) layout.Dimensions {
	return layout.Dimensions{Size: spacer}
}
