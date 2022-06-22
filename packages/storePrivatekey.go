package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func StorePrivateKey(publicKey string, privateKey string) string {

	//here we write a file.

	d1 := []byte("hello ")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	return "success"

}
