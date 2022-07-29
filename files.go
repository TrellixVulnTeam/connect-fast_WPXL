package main

import (
	"encoding/json"
	"fmt"

	"github.com/tse-lao/ipfs-user/ipfs"
)

func (a *App) IpfsRun() {
	ipfs.Init()
	/*
	   [] Make sure this is initialized correctly.

	*/

}
func (a *App) StartIPFS() interface{} {
	result := ipfs.RunDaemon()
	return result
}

func (a *App) StopIPFS() bool {
	ipfs.IpfsShutdown()

	return true
}

func (a *App) ReadDirectory(path string) []ipfs.FileStatus {
	result := ipfs.ReadDirectory(path)

	return result
}
func (a *App) FileStat() ipfs.FileStatus {
	result := ipfs.FileStat("/newProfile")
	fmt.Println(result)

	something := ipfs.FileStatus{}
	b, err := json.Marshal(something)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))

	return result
}
func (a *App) RemoveFile(path string) bool {
	status, _ := ipfs.RemoveFile(path)
	return status
}
func (a *App) CreateDirectory(name string) bool {
	status, message := ipfs.CreateFolder(name)
	fmt.Println(message)
	return status
}
func (a *App) ReadFile(path string) string {

	data := ipfs.ReadFile(path)

	return string(data)
}
func (a *App) AddFile(newPath, currentPath string) bool {
	status, message := ipfs.AddFile(newPath, currentPath)

	//here i need to cerate something that makes sure that
	fmt.Println(message)
	return status
}
