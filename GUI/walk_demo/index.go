package main

import (
	"strings"

	"github.com/lxn/walk"
	walkWin "github.com/lxn/walk/declarative"
)

func main() {
	var inTE, outTE *walk.TextEdit

	walkWin.MainWindow{
		Title:   "SCREAMO",
		MinSize: walkWin.Size{Width: 600, Height: 400},
		Layout:  walkWin.VBox{},
		Children: []walkWin.Widget{
			walkWin.HSplitter{
				Children: []walkWin.Widget{
					walkWin.TextEdit{AssignTo: &inTE},
					walkWin.TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			walkWin.PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}
