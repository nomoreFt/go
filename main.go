package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	srcDir := "D://test/moveDir"
	destDir := "D://test/dir"

	//c := make(chan string)

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		//go moveFile(file.Name(), srcDir, destDir, c)
		file.Name()
		copyFile(srcDir, destDir)

	}
	/*	for i := 0; i < len(files); i++ {
		result := <-c
		fmt.Println(result)
	}*/
}

func copyFile(srcDir, destDir string) {
	fin, err := os.Open(srcDir)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create(destDir)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}
func moveFile(fileNm string, srcDir string, destDir string, c chan string) {

	oldLocation := srcDir + "/" + fileNm
	newLocation := destDir + "/" + fileNm
	err := os.Rename(oldLocation, newLocation)

	if err != nil {
		log.Fatal(err)
	}
	c <- fileNm + "is moved"
}
