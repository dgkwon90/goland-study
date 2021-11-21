package image

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const RecommandText = "변환할 PNG 파일들을 이곳으로 드로그앤 드랍 해주세요!"

const StatusInit = "작업 대기 상태"
const StatusNullFile = "작업 파일 없음"
const StatusStart = "작업 진행중"
const StatusOK = "작업 완료"
const StatusNOK = "작업 실패"

type Options struct {
	Noise string
	Scale string
}

const WinWidth = 1000
const WinHeight = 800

func StartApp() {
	iconStop, err := walk.NewIconFromFile("./icon/stop.ico")
	if err != nil {
		log.Fatal(err)
	}
	iconWarning, err := walk.NewIconFromFile("./icon/warning.ico")
	if err != nil {
		log.Fatal(err)
	}
	iconStart, err := walk.NewIconFromFile("./icon/start.ico")
	if err != nil {
		log.Fatal(err)
	}
	iconOK, err := walk.NewIconFromFile("./icon/ok.ico")
	if err != nil {
		log.Fatal(err)
	}
	iconNOK, err := walk.NewIconFromFile("./icon/nok.ico")
	if err != nil {
		log.Fatal(err)
	}
	var textEdit *walk.TextEdit
	var statuBar *walk.StatusBarItem
	var inputFiles []string
	//fileCount := 0
	options := &Options{"3", "2"}

	MainWindow{
		Title:   "솔이를 위한 이미지 변환기 - v1.0",
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
			statuBar.SetText(StatusInit)
			statuBar.SetIcon(iconStop)
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
								Buttons: []RadioButton{
									RadioButton{
										Name:  "NoiseOption1",
										Text:  "없음",
										Value: "-1",
									},
									RadioButton{
										Name:  "NoiseOption2",
										Text:  "낮음",
										Value: "0",
									},
									RadioButton{
										Name:  "NoiseOption3",
										Text:  "중간",
										Value: "1",
									},
									RadioButton{
										Name:  "NoiseOption4",
										Text:  "높음",
										Value: "2",
									},
									RadioButton{
										Name:  "NoiseOption5",
										Text:  "최고",
										Value: "3",
									},
								},
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
								Buttons: []RadioButton{
									RadioButton{
										Name:  "ScaleOption1",
										Text:  "없음(크기 그대로)",
										Value: "1",
									},
									RadioButton{
										Name:  "ScaleOption2",
										Text:  "2x",
										Value: "2",
									},
									RadioButton{
										Name:  "ScaleOption3",
										Text:  "4x",
										Value: "4",
									},
									RadioButton{
										Name:  "ScaleOption4",
										Text:  "8x",
										Value: "5",
									},
									RadioButton{
										Name:  "ScaleOption5",
										Text:  "16x",
										Value: "16",
									},
									RadioButton{
										Name:  "ScaleOption6",
										Text:  "32x",
										Value: "32",
									},
								},
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
								statuBar.SetText(StatusStart + "(" + strconv.Itoa(successCount) + "/" + strconv.Itoa(fileCount) + ")")
								statuBar.SetIcon(iconStart)

								//run command
								//imageChangeCmd := exec.Command("go", "version")

								for i, filePath := range inputFiles {
									slice := strings.Split(filePath, "\\")
									fileName := slice[len(slice)-1]
									imageChangeCmd := exec.Command("./waifu2x-ncnn-vulkan-20210521-windows/waifu2x-ncnn-vulkan.exe",
										"-i", "\""+filePath+"\"",
										"-o", "./OutputImage/Change_"+fileName,
										"-n", options.Noise,
										"-s", options.Scale,
									)

									fmt.Println("Command : ", imageChangeCmd.String())

									if err := imageChangeCmd.Run(); err != nil {
										//if dateOut, err := imageChangeCmd.Output(); err != nil {
										//Fail
										//fmt.Println("Error: ", err)
										statuBar.SetText(StatusNOK)
										statuBar.SetIcon(iconNOK)
										inputFiles[i] += " => (Fail:" + err.Error() + ")"
									} else {
										//Success
										successCount++
										//fmt.Println("dateOut: ", string(dateOut))
										statuBar.SetText(StatusStart + "(" + strconv.Itoa(successCount) + "/" + strconv.Itoa(fileCount) + ")")
										inputFiles[i] += " => (Success)"
									}
									textEdit.SetText(strings.Join(inputFiles, "\r\n"))
								}
								if successCount == fileCount {
									statuBar.SetText(StatusOK + "(" + strconv.Itoa(fileCount) + "/" + strconv.Itoa(fileCount) + ")")
									statuBar.SetIcon(iconOK)
								}
							} else {
								//Warning
								statuBar.SetText(StatusNullFile)
								statuBar.SetIcon(iconWarning)
							}
						},
					},
				},
			},
		},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				AssignTo: &statuBar,
				Text:     StatusInit,
				Width:    WinWidth,
				Icon:     iconStop,
			},
		},
	}.Run()
}
