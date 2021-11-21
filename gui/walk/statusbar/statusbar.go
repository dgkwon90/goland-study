package statusbar

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func StartApp() {
	//icon1, err := walk.NewIconFromFile("../img/check.ico")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//icon2, err := walk.NewIconFromFile("../img/stop.ico")
	//if err != nil {
	//	log.Fatal(err)
	//}

	var sbi *walk.StatusBarItem

	MainWindow{
		Title:   "Walk Statusbar Example",
		MinSize: Size{600, 200},
		Layout:  VBox{MarginsZero: true},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				AssignTo: &sbi,
				//Icon:     icon1,
				Text:  "click",
				Width: 80,
				OnClicked: func() {
					if sbi.Text() == "click" {
						sbi.SetText("again")
						//sbi.SetIcon(icon2)
					} else {
						sbi.SetText("click")
						//sbi.SetIcon(icon1)
					}
				},
			},
			StatusBarItem{
				Text:        "left",
				ToolTipText: "no tooltip for me",
			},
			StatusBarItem{
				Text: "\tcenter",
			},
			StatusBarItem{
				Text: "\t\tright",
			},
			StatusBarItem{
				//Icon:        icon1,
				ToolTipText: "An icon with a tooltip",
			},
		},
	}.Run()
}
