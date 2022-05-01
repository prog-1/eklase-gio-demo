package screen

import (
	"eklase/state"
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// listStudents defines a screen layout for listing existing students.
func listStudents(th *material.Theme, vals *state.Handle) Screen {
	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return material.List(th, &list).Layout(gtx, vals.StudentCount(), func(gtx layout.Context, index int) layout.Dimensions {
					student := vals.Student(index)
					return material.Body1(th, fmt.Sprintf("%s %s", student.Surname, student.Name)).Layout(gtx)
				})
			}),
			layout.Rigid(material.Button(th, &close, "Close").Layout),
		)
		if close.Clicked() {
			return mainMenu(th, vals), d
		}
		return nil, d
	}
}
