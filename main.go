package main

import (
	"fmt"

	"github.com/lxn/walk"
	wui "github.com/lxn/walk/declarative"
	"github.com/ncruces/zenity"
)

func main() {
	var mainwin *walk.MainWindow
	var mi_open *walk.Action
	var s_status *walk.StatusBarItem
	wui.MainWindow{
		// Bounds: wui.Rectangle{
		// 	X:      100,
		// 	Y:      0,
		// 	Width:  500,
		// 	Height: 300,
		// },
		AssignTo: &mainwin,
		Name:     "demo",
		// Size: wui.Size{
		// 	Width:  500,
		// 	Height: 300,
		// },
		Title:  "demo",
		Layout: wui.VBox{},
		Children: []wui.Widget{
			wui.Label{
				Text: "ruok",
			},
			wui.PushButton{
				Text: "btn",
				OnClicked: func() {
					fmt.Println("click")
					var d, _ = walk.NewDialog(mainwin)
					d.SetTitle("ss")
					d.SetWidth(500)
					d.SetHeight(300)
					d.Show()
				},
			},
		},
		MenuItems: []wui.MenuItem{
			wui.Menu{
				Text: "file",
				Items: []wui.MenuItem{
					wui.Action{
						Text:     "open",
						AssignTo: &mi_open,
						OnTriggered: func() {
							fmt.Println("OnTriggered menu open")
						},
					},
				},
			},
		},
		StatusBarItems: []wui.StatusBarItem{
			{
				Text:     "status",
				AssignTo: &s_status,
				OnClicked: func() {
					var result = walk.MsgBox(mainwin, "tips", "status ok", walk.MsgBoxIconInformation|walk.MsgBoxYesNo)
					switch result {
					case walk.DlgCmdYes:
						fmt.Println("yes")
					case walk.DlgCmdNo:
						fmt.Println("ok")
					}
				},
			},
			{
				Text: "aaa",
				OnClicked: func() {
					zenity.Info("click aaa", zenity.Title("tip"), zenity.InfoIcon)
				},
			},
		},
		ContextMenuItems: []wui.MenuItem{
			wui.Action{
				Text: "exit app",
				OnTriggered: func() {
					mainwin.Close()
				},
			},
		},
	}.Run()
}
