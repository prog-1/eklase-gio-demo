package screen

import (
	"eklase/state"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// addStudent defines a screen layout for adding a new student.
func addStudent(th *material.Theme, state *state.Handle) Screen {
	var (
		name    widget.Editor
		surname widget.Editor

		close widget.Clickable
		save  widget.Clickable
	)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
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
			return mainMenu(th, state), d
		}
		if save.Clicked() {
			state.AddStudent(name.Text(), surname.Text())
			return mainMenu(th, state), d
		}
		return nil, d
	}
}
