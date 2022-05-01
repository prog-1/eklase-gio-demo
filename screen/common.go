package screen

import (
	"gioui.org/layout"
	"gioui.org/unit"
)

var in = layout.UniformInset(unit.Dp(5))

func rowInset(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { return in.Layout(gtx, w) }
}
func colInset(w layout.Widget) layout.Widget {
	return func(gtx layout.Context) layout.Dimensions { return in.Layout(gtx, w) }
}
