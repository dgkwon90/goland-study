package dropfiles

import (
	"fmt"
	. "github.com/lxn/walk/declarative"
	"golang.org/x/sys/windows"
	"os"
	"strings"
	"syscall"
	"time"
)

func runMeElevated() {
	verb := "runas"
	exe, _ := os.Executable()
	cwd, _ := os.Getwd()
	args := strings.Join(os.Args[1:], " ")

	verbPtr, _ := syscall.UTF16PtrFromString(verb)
	exePtr, _ := syscall.UTF16PtrFromString(exe)
	cwdPtr, _ := syscall.UTF16PtrFromString(cwd)
	argPtr, _ := syscall.UTF16PtrFromString(args)

	var showCmd int32 = 1 //SW_NORMAL

	err := windows.ShellExecute(0, verbPtr, exePtr, argPtr, cwdPtr, showCmd)
	if err != nil {
		fmt.Println(err)
	}
}

func amAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	if err != nil {
		fmt.Println("admin no")
		return false
	}
	fmt.Println("admin yes")
	return true
}

func StartApp() {
	//var textEdit *walk.TextEdit

	if !amAdmin() {
		runMeElevated()
	}
	time.Sleep(10 * time.Second)

	MainWindow{
		Title:   "Walk DropFiles Example",
		MinSize: Size{320, 240},
		Layout:  VBox{},
		OnDropFiles: func(files []string) {
			fmt.Println("OnDrop!!!")
			//textEdit.SetText(strings.Join(files, "\r\n"))
		},
		//Children: []Widget{
		//	TextEdit{
		//		AssignTo: &textEdit,
		//		ReadOnly: true,
		//		Text:     "Drop files here, from windows explorer...",
		//	},
		//},
	}.Run()
}
