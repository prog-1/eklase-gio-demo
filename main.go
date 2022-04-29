package main

import (
	"log"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/jmoiron/sqlx"

	_ "modernc.org/sqlite"
)

var (
	createTableStmt = `
	CREATE TABLE students (
		id	INTEGER,
		name	TEXT,
		surname	TEXT,
		PRIMARY KEY(id AUTOINCREMENT)
	);
`
)

func main() {
	db := sqlx.MustOpen("sqlite", ":memory:")
	defer db.Close()
	db.MustExec(createTableStmt)

	go func() {
		w := app.NewWindow()
		err := draw(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	var ops op.Ops
	th := material.NewTheme(gofont.Collection())
	var addStudentBtn, listStudentsBtn, quitBtn widget.Clickable
	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)
			drawMenu(gtx, th, &addStudentBtn, &quitBtn, &listStudentsBtn)
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			return e.Err

		}
	}
	return nil
}

func drawMenu(
	gtx layout.Context,
	th *material.Theme,
	addStudentBtn, listStudentsBtn, quitBtn *widget.Clickable) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx,
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, addStudentBtn, "Add Student")
				return btn.Layout(gtx)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, listStudentsBtn, "List Students")
				return btn.Layout(gtx)
			},
		),
		layout.Rigid(
			func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(th, quitBtn, "Quit")
				return btn.Layout(gtx)
			},
		),
	)
}

// db := sqlx.MustOpen("sqlite", ":memory:")
// defer db.Close()
// db.MustExec(createTableStmt)
// for _, name := range []string{"foo", "bar", "larch", "spam", "egg"} {
// 	db.MustExec(insertStmt, name)
// }
// var entries []struct {
// 	ID   int64  `db:"id"`
// 	Name string `db:"name"`
// }
// db.Select(&entries, selectStmt)
// fmt.Println(entries)
