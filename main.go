package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	targetDir := "D://test/moveDir"
	newDir := "D://test/dir"

	c := make(chan string)

	files, err := ioutil.ReadDir(targetDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		go moveFile(file.Name(), targetDir, newDir, c)

	}
	for i := 0; i < len(files); i++ {
		result := <-c
		fmt.Println(result)
	}

}

func moveFile(fileNm string, targetDir string, newDir string, c chan string) {

	oldLocation := targetDir + "/" + fileNm
	newLocation := newDir + "/" + fileNm
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
	c <- fileNm + "is moved"
}
