package image

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

const WinWidth = 1000
const WinHeight = 800

const RecommandText = "변환할 PNG 파일들을 이곳으로 드로그앤 드랍 해주세요!"

const StatusInit = "작업 대기 상태"
const StatusNullFindPath = "경로가 비어 있음"
const StatusNullFile = "작업 파일 없음"
const StatusStart = "작업 진행중"
const StatusOK = "작업 완료"
const StatusNOK = "작업 실패"

type Options struct {
	Noise string
	Scale string
}

var iconStop, iconWarning, iconStart, iconOK, iconNOK *walk.Icon

func initStatusIcon() {
	iconStop, _ = walk.NewIconFromFile("./icon/stop.ico")
	iconWarning, _ = walk.NewIconFromFile("./icon/warning.ico")
	iconStart, _ = walk.NewIconFromFile("./icon/start.ico")
	iconOK, _ = walk.NewIconFromFile("./icon/ok.ico")
	iconNOK, _ = walk.NewIconFromFile("./icon/nok.ico")
}

func StartApp() {
	//var findPathTextEdit *walk.TextEdit
	var textEdit *walk.TextEdit
	var statuBar *walk.StatusBarItem
	initStatusIcon()
	var inputFiles []string
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
						//Children: []Widget{
						//	Label{
						//		Text: "변경 Image 경로 :",
						//	},
						//	TextEdit {
						//		AssignTo: &findPathTextEdit,
						//		CompactHeight:true,
						//	},
						//	PushButton{
						//		Text: "검색",
						//		OnClicked: func() {
						//			if len(findPathTextEdit.Text()) <= 0 {
						//				//warning
						//				statuBar.SetText(StatusNullFile)
						//				statuBar.SetIcon(iconWarning)
						//			} else {
						//				var findFiles []string
						//				err := filepath.Walk(findPathTextEdit.Text(), func(path string, info os.FileInfo, err error) error {
						//					if err != nil {
						//						fmt.Println(err)
						//						return nil
						//					}
						//					if !info.IsDir() &&
						//						(filepath.Ext(path) == ".png"  ||
						//							(filepath.Ext(path) == ".PNG") ||
						//							(filepath.Ext(path) == ".jpg") ||
						//							(filepath.Ext(path) == ".JPG") ||
						//							(filepath.Ext(path) == ".jpeg") ||
						//							(filepath.Ext(path) == ".JPEG")) {
						//						newPath := strings.Replace(path, "\\", "/", -1)
						//						findFiles = append(findFiles, newPath)
						//					}
						//					return nil
						//				})
						//				if err != nil || len(findFiles) <= 0 {
						//					//Warning
						//					statuBar.SetText(StatusNullFile)
						//					statuBar.SetIcon(iconWarning)
						//				} else {
						//					inputFiles = findFiles
						//					fmt.Println("Drop count : ", len(inputFiles))
						//					textEdit.SetText(strings.Join(findFiles, "\r\n"))
						//					statuBar.SetText(StatusInit)
						//					statuBar.SetIcon(iconStop)
						//				}
						//			}
						//		},
						//	},
						//},
					},

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
								resultFiles := inputFiles
								statuBar.SetText(StatusStart + "(" + strconv.Itoa(successCount) + "/" + strconv.Itoa(fileCount) + ")")
								statuBar.SetIcon(iconStart)

								for i, filePath := range inputFiles {

									newPath := strings.Replace(filePath, "\\", "/", -1)
									slice := strings.Split(newPath, "/")
									fileName := slice[len(slice)-1]

									imageChangeCmd := exec.Command("./waifu2x-ncnn-vulkan-20210521-windows/waifu2x-ncnn-vulkan.exe",
										"-i", newPath,
										"-o", "./OutputImage/Change_"+fileName,
										"-n", options.Noise,
										"-s", options.Scale,
									)

									//imageChangeCmd := exec.Command("./waifu2x-ncnn-vulkan-20210521-windows/waifu2x-ncnn-vulkan.exe",
									//	"-i", "E:/Workspace/Go/src/04_dgkwon90/InputImage/1.jpg",
									//	"-o", "./OutputImage/Change_1.jpg",
									//	"-n", "3",
									//	"-s", "2",
									//)
									//
									//if err := imageChangeCmd.Run(); err != nil {
									//	//Fail
									//	fmt.Println("Error: ", err)
									//} else {
									//	//Success
									//	fmt.Println(imageChangeCmd.String())
									//	fmt.Println("Success")
									//}

									imageChangeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
									fmt.Println("Command : \n", imageChangeCmd.String())

									if err := imageChangeCmd.Run(); err != nil {
										//Fail
										fmt.Println("Error: ", err)
										statuBar.SetText(StatusNOK)
										statuBar.SetIcon(iconNOK)
										resultFiles[i] += " => (Fail:" + err.Error() + ")"
									} else {
										//Success
										successCount++
										statuBar.SetText(StatusStart + "(" + strconv.Itoa(successCount) + "/" + strconv.Itoa(fileCount) + ")")
										resultFiles[i] += " => (Success)"
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
