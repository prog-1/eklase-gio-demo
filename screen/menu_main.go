package screen

import (
	"eklase/manager"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// mainMenu defines a main menu screen layout.
func mainMenu(th *material.Theme, manager *manager.AppManager) Screen {
	var (
		add  widget.Clickable
		list widget.Clickable
		quit widget.Clickable
	)
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(material.Button(th, &add, "Add student").Layout)),
			layout.Rigid(rowInset(material.Button(th, &list, "List students").Layout)),
			layout.Rigid(rowInset(material.Button(th, &quit, "Quit").Layout)),
		)
		if add.Clicked() {
			return addStudent(th, manager), d
		}
		if list.Clicked() {
			return listStudents(th, manager), d
		}
		if quit.Clicked() {
			manager.Quit()
		}
		return nil, d
	}
}
