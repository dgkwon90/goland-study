package example

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func StartApp() {
	var inTE, outTE *walk.TextEdit

	MainWindow{
		Title:   "SCREAMO",
		MinSize: Size{400, 200},
		Layout:  VBox{},
		Children: []Widget{
			HSplitter{
				Children: []Widget{
					TextEdit{AssignTo: &inTE},
					TextEdit{AssignTo: &outTE, ReadOnly: true},
				},
			},
			PushButton{
				Text: "SCREAM",
				OnClicked: func() {
					outTE.SetText(strings.ToUpper(inTE.Text()))
				},
			},
		},
	}.Run()
}

func StartApp2() {
	var te *walk.TextEdit

	if _, err := (MainWindow{
		Title:   "Walk Clipboard Example",
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			PushButton{
				Text: "Copy",
				OnClicked: func() {
					if err := walk.Clipboard().SetText(te.Text()); err != nil {
						log.Print("Copy: ", err)
					}
				},
			},
			PushButton{
				Text: "Paste",
				OnClicked: func() {
					if text, err := walk.Clipboard().Text(); err != nil {
						log.Print("Paste: ", err)
					} else {
						te.SetText(text)
					}
				},
			},
			TextEdit{
				AssignTo: &te,
			},
		},
	}).Run(); err != nil {
		log.Fatal(err)
	}
}

func StartApp3() {
	var files []string
	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {

			fmt.Println(err)
			return nil
		}

		if !info.IsDir() && filepath.Ext(path) == ".txt" {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}
