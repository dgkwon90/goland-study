package window

import (
	"fmt"
	"github.com/dgkwon90/golang-study/moon-tools/image-upscale/controller"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strconv"
	"strings"
)

const WinWidth = 1000
const WinHeight = 800

const TitleText = "Moon님을 위한 이미지 스케일 업 변환기 - MoonUpScale v2.0"

const RecommandText = "변환할 PNG 파일들을 이곳으로 드로그앤 드랍 해주세요!"

const StatusInit = "작업 대기 상태"
const StatusNullFile = "작업 파일 없음"
const StatusStart = "작업 진행중"
const StatusOK = "작업 완료"
const StatusNOK = "작업 실패"

var mainWindow MainWindow

var iconTitle, iconStop, iconWarning, iconStart, iconOK, iconNOK *walk.Icon

func initStatusIcon() {
	iconTitle, _ = walk.NewIconFromFile("./asset/icon/MoonUpScale.ico")
	iconStop, _ = walk.NewIconFromFile("./asset/icon/stop.ico")
	iconWarning, _ = walk.NewIconFromFile("./asset/icon/warning.ico")
	iconStart, _ = walk.NewIconFromFile("./asset/icon/start.ico")
	iconOK, _ = walk.NewIconFromFile("./asset/icon/ok.ico")
	iconNOK, _ = walk.NewIconFromFile("./asset/icon/nok.ico")
}

type options struct {
	Noise string
	Scale string
}

func addNoiseRadioButtonList() (noiseRadioButton []RadioButton) {
	noiseRadioButton = []RadioButton{
		RadioButton{
			Name: "NoiseOption1", Text: "없음", Value: "-1",
		},
		RadioButton{
			Name: "NoiseOption2", Text: "낮음", Value: "0",
		},
		RadioButton{
			Name: "NoiseOption3", Text: "중간", Value: "1",
		},
		RadioButton{
			Name: "NoiseOption4", Text: "높음", Value: "2",
		},
		RadioButton{
			Name: "NoiseOption5", Text: "최고", Value: "3",
		},
	}
	return noiseRadioButton
}

func addScaleRadioButtonList() (scaleRadioButton []RadioButton) {
	scaleRadioButton = []RadioButton{
		RadioButton{
			Name: "ScaleOption1", Text: "없음(크기 그대로)", Value: "1",
		},
		RadioButton{
			Name: "ScaleOption2", Text: "2x", Value: "2",
		},
		RadioButton{
			Name: "ScaleOption3", Text: "4x", Value: "4",
		},
		RadioButton{
			Name: "ScaleOption4", Text: "8x", Value: "5",
		},
		RadioButton{
			Name: "ScaleOption5", Text: "16x", Value: "16",
		},
		RadioButton{
			Name: "ScaleOption6", Text: "32x", Value: "32",
		},
	}
	return scaleRadioButton
}

func Load() {
	initStatusIcon()
	var textEdit *walk.TextEdit
	var statusBar *walk.StatusBarItem

	var inputFiles []string
	options := &options{"3", "2"}

	mainWindow = MainWindow{
		Title:   TitleText,
		Icon:    iconTitle,
		Size:    Size{WinWidth, WinHeight},
		MinSize: Size{WinWidth, WinHeight},
		Layout:  VBox{},
		DataBinder: DataBinder{
			DataSource: options,
			AutoSubmit: true,
			OnSubmitted: func() {
				fmt.Println(options)
			},
		},
		OnDropFiles: func(files []string) {
			inputFiles = files
			fmt.Println("Drop count : ", len(inputFiles))
			textEdit.SetText(strings.Join(files, "\r\n"))
			statusBar.SetText(StatusInit)
			statusBar.SetIcon(iconStop)
		},
		Children: []Widget{
			VSplitter{
				Children: []Widget{
					HSplitter{
						Children: []Widget{
							Label{
								Text: "노이크 감소 :",
							},
							RadioButtonGroup{
								DataMember: "Noise",
								Buttons:    addNoiseRadioButtonList(),
							},
						},
					},
					HSplitter{
						Children: []Widget{
							Label{
								Text: "확대 :",
							},
							RadioButtonGroup{
								DataMember: "Scale",
								Buttons:    addScaleRadioButtonList(),
							},
						},
					},
					TextEdit{
						AssignTo: &textEdit,
						ReadOnly: true,
						Text:     RecommandText,
					},
					PushButton{
						Text: "변환",
						OnClicked: func() {
							fileCount := len(inputFiles)
							fmt.Println("Drop count : ", fileCount)
							fmt.Println("Options Noise : ", options.Noise)
							fmt.Println("Options Scale : ", options.Scale)
							if fileCount > 0 {
								//start
								successCount := 0
								resultFiles := inputFiles
								statusBar.SetText(StatusStart + "(" + strconv.Itoa(successCount) + "/" + strconv.Itoa(fileCount) + ")")
								statusBar.SetIcon(iconStart)
								for i, filePath := range inputFiles {
									newPath := strings.Replace(filePath, "\\", "/", -1)
									if err := controller.RunImageParse(newPath, options.Noise, options.Scale); err != nil {
										//Fail
										fmt.Println("Error: ", err)
										statusBar.SetText(StatusNOK)
										statusBar.SetIcon(iconNOK)
										resultFiles[i] += " => (Fail:" + err.Error() + ")"
									} else {
										//Success
										successCount++
										statusBar.SetText(StatusStart + "(" + strconv.Itoa(successCount) + "/" + strconv.Itoa(fileCount) + ")")
										resultFiles[i] += " => (Success)"
									}
									textEdit.SetText(strings.Join(inputFiles, "\r\n"))
								}
								if successCount == fileCount {
									statusBar.SetText(StatusOK + "(" + strconv.Itoa(fileCount) + "/" + strconv.Itoa(fileCount) + ")")
									statusBar.SetIcon(iconOK)
								}
							} else {
								//Warning
								statusBar.SetText(StatusNullFile)
								statusBar.SetIcon(iconWarning)
							}
						},
					},
				},
			},
		},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				AssignTo: &statusBar,
				Width:    WinWidth,
				Text:     StatusInit,
				Icon:     iconStop,
			},
		},
	}
}

func Start() {
	mainWindow.Run()
}
