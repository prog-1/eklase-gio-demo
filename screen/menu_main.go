package screen

import (
	"eklase/state"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// mainMenu defines a main menu screen layout.
func mainMenu(th *material.Theme, state *state.Handle) Screen {
	var (
		add  widget.Clickable
		list widget.Clickable
		quit widget.Clickable
	)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(material.Button(th, &add, "Add student").Layout),
			layout.Rigid(material.Button(th, &list, "List students").Layout),
			layout.Rigid(material.Button(th, &quit, "Quit").Layout),
		)
		if add.Clicked() {
			return addStudent(th, state), d
		}
		if list.Clicked() {
			return listStudents(th, state), d
		}
		if quit.Clicked() {
			state.Quit()
		}
		return nil, d
	}
}
