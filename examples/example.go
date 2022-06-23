package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	customFileExplorer "github.com/nomoreFt/woodieFileControl"
)

func main() {
	router := httprouter.New()
	//outer.GET("/", Index)
	router.GET("/copyDir", copyFile)
	router.GET("/moveDir", Index)

	log.Fatal(http.ListenAndServe(":8080", router))

	nouse := "test"

	fmt.Print("응 못해" + nouse)

}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	srcDir := "D://test/srcDir"
	destDir := "C://test/destDir"

	fmt.Fprintf(w, srcDir+"에서 "+destDir+"으로 파일 Copy가 진행중입니다.")
	c := make(chan string)

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		fmt.Println(w, err.Error())
	}
	for _, file := range files {
		go customFileExplorer.MoveFile(file.Name(), srcDir, destDir, c)
		file.Name()
		//go customFileExplorer.CopyFile(srcDir+"/"+file.Name(), destDir+"/"+file.Name(), c)

	}
	for i := 0; i < len(files); i++ {
		result := <-c
		fmt.Println(w, result)
	}
}
func copyFile(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	srcDir := "D://test/srcDir"
	destDir := "D://test/destDir"

	fmt.Fprintf(w, srcDir+"에서 "+destDir+"으로 파일 Copy가 진행중입니다.")
	c := make(chan string)

	files, err := ioutil.ReadDir(srcDir)
	if err != nil {
		fmt.Println(w, err.Error())
	}
	for _, file := range files {
		//go customFileExplorer.MoveFile(file.Name(), srcDir, destDir, c)
		file.Name()
		go customFileExplorer.CopyFile(srcDir+"/"+file.Name(), destDir+"/"+file.Name(), c)

	}
	for i := 0; i < len(files); i++ {
		result := <-c
		fmt.Println(w, result)
	}
}
