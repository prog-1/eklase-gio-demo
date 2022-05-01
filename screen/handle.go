package screen

import (
	"eklase/manager"
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
func NewHandle(manager *manager.AppManager) (*Handle, error) {
	th := material.NewTheme(gofont.Collection())
	if th == nil {
		return nil, errors.New("unexpected error while loading theme")
	}

	h := Handle{
		window: app.NewWindow(),
		theme:  th,
		// mainMenus is the default page.
		layout: mainMenu(th, manager),
	}
	return &h, nil
}

// HandleEvents handles application events.
func (h *Handle) HandleEvents(manager *manager.AppManager) error {
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
			if manager.ShouldQuit() {
				h.window.Perform(system.ActionClose)
			}
			evt.Frame(gtx.Ops)
		case system.DestroyEvent:
			return evt.Err
		}
	}
	return nil
}
