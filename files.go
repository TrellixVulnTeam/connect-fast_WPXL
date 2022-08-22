package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	crypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/tse-lao/ether-user/wallet"
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
	fmt.Println(data)
	return string(data)
}

func (a *App) GetIPFS(path string) string {
	bool, data := ipfs.DownloadFromIPFS(path)
	fmt.Println(bool, data)
	//read the file and delete the fie.
	dat, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return "There is an error occured"
	}

	//remove the file from the path.
	os.Remove(path)

	return string(dat)

}

func (a *App) WriteFile(data, extension string) string {
	newFile, _ := os.Create("temporaryFile." + extension)

	newFile.Write([]byte(data))

	//current path
	path, _ := os.Getwd()

	fmt.Println(path + "/temporaryFile." + extension)

	status, result := ipfs.UploadToIPFS(path + "/temporaryFile." + extension)

	s := strings.Split(result, " ")

	fmt.Println(status)
	return s[1]
}

func (a *App) AddFile(newPath, currentPath string) bool {
	status, message := ipfs.AddFile(newPath, currentPath)

	//here i need to cerate something that makes sure that
	fmt.Println(message)
	return status
}

func (a *App) EncryptToSmart(publicKey, data string) string {
	fmt.Println(publicKey)
	fmt.Println(data)

	//public key change

	PublicKeyToBytes, error := hexutil.Decode(publicKey)

	if error != nil {
		fmt.Println(error)
	}
	realPublicKey, err := crypto.UnmarshalPubkey(PublicKeyToBytes)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(realPublicKey)
	newData := wallet.EncryptWithPublicKey(realPublicKey, []byte(data))

	newFile, _ := os.Create("temporaryFile")

	newFile.Write(newData)

	//current path
	path, _ := os.Getwd()

	fmt.Println(path + "/temporaryFile.txt")

	status, result := ipfs.UploadToIPFS(path + "/temporaryFile.txt")

	s := strings.Split(result, " ")

	fmt.Println(status)
	return s[1]
}

func (a *App) DecryptData(path string) string {
	bool, data := ipfs.DownloadFromIPFS(path)
	fmt.Println(bool, data)
	//read the file and delete the fie.
	dat, err := os.ReadFile(path)

	if err != nil {
		fmt.Println(err)
		return "There is an error occured"
	}

	result := wallet.DecryptWithPrivateKey("password", dat)

	fmt.Println(string(result))

	return string(result)
}

/* c */
