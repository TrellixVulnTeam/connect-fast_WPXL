package main

import (
	"fmt"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileStatus struct {
	Name         string      `json:"name"`
	Path         string      `json:"path"`
	Size         int64       `json:"size"`
	LastModified string      `json:"last_modified"`
	Dir          bool        `json:"is_dir"`
	Sys          interface{} `json:"sys"`
}

func (a *App) SelectFile() FileStatus {
	selection, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select File",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			},
			{
				DisplayName: "Videos (*.mov;*.mp4)",
				Pattern:     "*.mov;*.mp4",
			},
		},
	})

	if err != nil {
		fmt.Println("Error")

	}
	fmt.Println(selection)
	file := checkFileStatus(selection)
	return file
}

func checkFileStatus(path string) FileStatus {
	fileStat, err := os.Stat(path)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileStat)

	file := FileStatus{}
	file.Path = path
	//make sure it is in MB
	file.Name = fileStat.Name()
	file.Size = fileStat.Size()
	file.Sys = fileStat.Sys()
	file.LastModified = fileStat.ModTime().String()
	file.Dir = fileStat.IsDir()

	return file
}
