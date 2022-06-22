package woodieFileExplorer

import (
	"io"
	"log"
	"os"
)

func CopyFile(srcFileNm, destFileNm string, c chan string) {
	fin, err := os.Open(srcFileNm)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create(destFileNm)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
	c <- srcFileNm + " is copied"
}
func MoveFile(fileNm string, srcDir string, destDir string, c chan string) {

	oldLocation := srcDir + "/" + fileNm
	newLocation := destDir + "/" + fileNm
	err := os.Rename(oldLocation, newLocation)

	if err != nil {
		log.Fatal(err)
	}
	c <- fileNm + " is moved"
}
