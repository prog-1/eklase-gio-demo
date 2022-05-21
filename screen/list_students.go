package screen

import (
	"eklase/state"
	"image"
	"image/color"
	"log"
	"strconv"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// ListStudent defines a screen layout for listing existing students.
func ListStudent(th *material.Theme, state *state.State) Screen {
	var close widget.Clickable
	list := widget.List{List: layout.List{Axis: layout.Vertical}}

	lightContrast := th.ContrastBg
	lightContrast.A = 0x11
	darkContrast := th.ContrastBg
	darkContrast.A = 0x33

	students, err := state.Students()
	if err != nil {
		// TODO: Show user an error toast.
		log.Printf("failed to fetch students: %v", err)
		return nil
	}

	layoutCell := func(gtx layout.Context, bg color.NRGBA, text string) layout.Dimensions {
		return layout.Stack{}.Layout(gtx,
			layout.Expanded(func(gtx layout.Context) layout.Dimensions {
				max := image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Min.Y)
				paint.FillShape(gtx.Ops, bg, clip.Rect{Max: max}.Op())
				return layout.Dimensions{Size: gtx.Constraints.Min}
			}),
			layout.Stacked(rowInset(material.Body1(th, text).Layout)),
		)
	}
	layoutRow := func(gtx layout.Context, bg color.NRGBA, id, last, first string) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(2, func(gtx layout.Context) layout.Dimensions {
				return layoutCell(gtx, bg, id)
			}),
			layout.Flexed(5, func(gtx layout.Context) layout.Dimensions {
				return layoutCell(gtx, bg, last)
			}),
			layout.Flexed(5, func(gtx layout.Context) layout.Dimensions {
				return layoutCell(gtx, bg, first)
			}),
		)
	}
	studentsLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return layoutRow(gtx, th.Bg, "ID", "Last Name", "First Name")
			}),
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				return material.List(th, &list).Layout(gtx, len(students), func(gtx layout.Context, index int) layout.Dimensions {
					student := students[index]
					bg := lightContrast
					if index%2 == 0 {
						bg = darkContrast
					}
					return layoutRow(gtx, bg, strconv.FormatInt(student.ID, 10), student.Surname, student.Name)
				})
			}),
		)
	}

	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Flexed(1, rowInset(studentsLayout)),
			layout.Rigid(rowInset(material.Button(th, &close, "Close").Layout)),
		)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		return nil, d
	}
}
