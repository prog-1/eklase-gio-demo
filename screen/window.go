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

// Window is a UI handler.
type Window struct {
	// window defines a window of the GIO framework.
	window *app.Window
	// theme is a UI theme.
	theme *material.Theme
	// layout defines a screen layout to render.
	layout Screen
}

// Screen defines the current layout.
type Screen func(gtx layout.Context) (Screen, layout.Dimensions)

// NewWindow creates new Window.
func NewWindow(state *state.State) (*Window, error) {
	th := material.NewTheme(gofont.Collection())
	if th == nil {
		return nil, errors.New("unexpected error while loading theme")
	}

	h := Window{
		window: app.NewWindow(),
		theme:  th,
		// mainMenus is the default page.
		layout: mainMenu(th, state),
	}
	return &h, nil
}

// HandleEvents handles application events.
func (w *Window) HandleEvents(state *state.State) error {
	for e := range w.window.Events() {
		switch evt := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&op.Ops{}, evt)
			layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				nextLayout, d := w.layout(gtx)
				if nextLayout != nil {
					w.layout = nextLayout
				}
				return d
			})
			if state.ShouldQuit() {
				w.window.Perform(system.ActionClose)
			}
			evt.Frame(gtx.Ops)
		case system.DestroyEvent:
			return evt.Err
		}
	}
	return nil
}
