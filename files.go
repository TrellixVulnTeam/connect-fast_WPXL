package main

import (
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

func (a *App) ReadDirectory(path string) interface{} {
	result := ipfs.ReadDirectory(path)
	fmt.Println(result)
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
func (a *App) ReadFile(path string) []byte {
	return ipfs.ReadFile(path)
}
func (a *App) AddFile(currentPath string, newPath string) bool {
	status, _ := ipfs.AddFile(currentPath, newPath)

	return status
}
