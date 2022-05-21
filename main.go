package main

import (
	"log"
	"os"

	"eklase/screen"
	"eklase/state"
	"eklase/storage"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"

	_ "modernc.org/sqlite"
)

func main() {

	// Run the main event loop.
	go func() {
		w := app.NewWindow(app.Title("e-Klasse"))
		if err := mainLoop(w); err != nil {
			log.Fatalf("failed to handle events: %v", err)
		}
		// Gracefully exit the application at the end.
		os.Exit(0)
	}()
	app.Main()
}

func mainLoop(w *app.Window) error {
	stor, err := storage.Open("school.db")
	if err != nil {
		return err
	}
	defer stor.Close()

	appState := state.New(stor)

	th := material.NewTheme(gofont.Collection())
	currentLayout := screen.MainMenu(th, appState)

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(&op.Ops{}, e)
				layout.UniformInset(unit.Dp(5)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					nextLayout, d := currentLayout(gtx)
					if nextLayout != nil {
						currentLayout = nextLayout
					}
					return d
				})
				if appState.ShouldQuit() {
					w.Perform(system.ActionClose)
				}
				e.Frame(gtx.Ops)
			case system.DestroyEvent:
				return e.Err
			}
		}
	}
}
