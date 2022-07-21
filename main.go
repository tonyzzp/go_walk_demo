package main

import (
	"fmt"

	wui "github.com/lxn/walk/declarative"
)

func main() {
	wui.MainWindow{
		Name: "demo",
		Size: wui.Size{
			Width:  500,
			Height: 300,
		},
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
				},
			},
		},
	}.Run()
}
