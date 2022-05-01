// Package ui contains logic for application graphic rendering.
package ui

import (
	"eklase/state"
	"errors"
	"fmt"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
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
	layout screen
}

// screen defines the current layout.
// A function that renders a page is expected to return screen, i.e.
//
// return func(gtx layout.Context) screen {
//   layout.Flex{...}.Layout(gtx,
//     ...
//   )
//   // Handle button clicks and other events here.
// }
type screen func(gtx layout.Context) screen

// New initializes the UI handler.
func New(state *state.Handle) (*Handle, error) {
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
func (h *Handle) Close() error {
	return h.state.Close()
}

// HandleEvents handles application events.
func (h *Handle) HandleEvents() error {
	for e := range h.window.Events() {
		switch evt := e.(type) {
		case system.FrameEvent:
			h.displayWindow(evt)
		case system.DestroyEvent:
			return evt.Err
		}
	}
	return nil
}

// displayWindow renders a page layout and handles application events.
func (h *Handle) displayWindow(evt system.FrameEvent) {
	gtx := layout.NewContext(&op.Ops{}, evt)
	if nextLayout := h.layout(gtx); nextLayout != nil {
		h.layout = nextLayout
	}
	if h.state.ShouldQuit() {
		h.window.Perform(system.ActionClose)
	}
	evt.Frame(gtx.Ops)
}

// mainMenu defines a main menu page layout.
func mainMenu(th *material.Theme, state *state.Handle) screen {
	var (
		add  widget.Clickable
		list widget.Clickable
		quit widget.Clickable
	)
	return func(gtx layout.Context) screen {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(material.Button(th, &add, "Add student").Layout),
			layout.Rigid(material.Button(th, &list, "List students").Layout),
			layout.Rigid(material.Button(th, &quit, "Quit").Layout),
		)
		if add.Clicked() {
			return addStudent(th, state)
		}
		if list.Clicked() {
			return listStudents(th, state)
		}
		if quit.Clicked() {
			state.Quit()
		}
		return nil
	}
}

// addStudent defines a page layout for adding a new student.
func addStudent(th *material.Theme, state *state.Handle) screen {
	var (
		name    widget.Editor
		surname widget.Editor

		close widget.Clickable
		save  widget.Clickable
	)
	return func(gtx layout.Context) screen {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Flexed(1, material.Editor(th, &name, "First name").Layout),
					layout.Flexed(1, material.Editor(th, &surname, "Last name").Layout),
				)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
					layout.Flexed(1, material.Button(th, &close, "Close").Layout),
					layout.Flexed(1, material.Button(th, &save, "Save").Layout),
				)
			}),
		)
		if close.Clicked() {
			return mainMenu(th, state)
		}
		if save.Clicked() {
			state.AddStudent(name.Text(), surname.Text())
			return mainMenu(th, state)
		}
		return nil
	}
}

// listStudents defines a page layout for listing existing students.
func listStudents(th *material.Theme, vals *state.Handle) screen {
	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}
	return func(gtx layout.Context) screen {
		layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return material.List(th, &list).Layout(gtx, vals.StudentCount(), func(gtx layout.Context, index int) layout.Dimensions {
					student := vals.Student(index)
					return material.Body1(th, fmt.Sprintf("%s %s", student.Surname, student.Name)).Layout(gtx)
				})
			}),
			layout.Rigid(material.Button(th, &close, "Close").Layout),
		)
		if close.Clicked() {
			return mainMenu(th, vals)
		}
		return nil
	}
}
