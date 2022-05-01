package screen

import (
	"eklase/state"
	"errors"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

// Handle is a UI handler.
type Handle struct {
	// window defines a window of the GIO framework.
	window *app.Window
	// theme is a UI theme.
	theme *material.Theme
	// state is an application state.
	state *state.Handle
	// layout defines a screen layout to render.
	layout Screen
}

// Screen defines the current layout.
// A function that renders a page is expected to return Screen, i.e.
//
// return func(gtx layout.Context) Screen {
//   layout.Flex{...}.Layout(gtx,
//     ...
//   )
//   // Handle button clicks and other events here.
// }
type Screen func(gtx layout.Context) (Screen, layout.Dimensions)

// NewHandle initializes the UI handler.
func NewHandle(state *state.Handle) (*Handle, error) {
	th := material.NewTheme(gofont.Collection())
	if th == nil {
		return nil, errors.New("unexpected error while loading theme")
	}

	h := Handle{
		window: app.NewWindow(),
		theme:  th,
		state:  state,
		// mainMenus is the default page.
		layout: mainMenu(th, state),
	}
	return &h, nil
}

// Close closes the UI dependencies.
func (h *Handle) Close() error { return h.state.Close() }

// HandleEvents handles application events.
func (h *Handle) HandleEvents() error {
	for e := range h.window.Events() {
		switch evt := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&op.Ops{}, evt)
			layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				nextLayout, d := h.layout(gtx)
				if nextLayout != nil {
					h.layout = nextLayout
				}
				return d
			})
			if h.state.ShouldQuit() {
				h.window.Perform(system.ActionClose)
			}
			evt.Frame(gtx.Ops)
		case system.DestroyEvent:
			return evt.Err
		}
	}
	return nil
}
