package controller

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
)

func RunImageParse(imagePath, noise, scale string) error {
	slice := strings.Split(imagePath, "/")
	fileName := slice[len(slice)-1]
	imageChangeCmd := exec.Command("./lib/waifu2x-ncnn-vulkan-20210521-windows/waifu2x-ncnn-vulkan.exe",
		"-i", imagePath,
		"-o", "./OutputImage/Change_"+fileName,
		"-n", noise,
		"-s", scale,
	)

	imageChangeCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	fmt.Println("Command : \n", imageChangeCmd.String())
	return imageChangeCmd.Run()
}
