package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	srcDir := "D://test/srcDir"
	destDir := "D://test/destDir"

	c := make(chan string)

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		fmt.Println(err)
	}
	for _, file := range files {
		//go moveFile(file.Name(), srcDir, destDir, c)
		file.Name()
		go copyFile(srcDir+"/"+file.Name(), destDir+"/"+file.Name(), c)

	}
	for i := 0; i < len(files); i++ {
		result := <-c
		fmt.Println(result)
	}
}
