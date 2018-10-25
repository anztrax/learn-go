package filehandling

import (
	"fmt"
	"github.com/gobuffalo/packr"
	"io/ioutil"
	"os"
	"path/filepath"
)

func getExecutablePath() string{
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

func fileHandlingUsingBox(){
	box := packr.NewBox("./")
	dataBox := box.String("test.txt")
	fmt.Println("contents of file : ", dataBox, "\n")

}

func TryFileHandling(){
	fmt.Printf("\nexecutable path : %s\n\n", getExecutablePath())
	data , err := ioutil.ReadFile("/Users/andrewananta/go/src/testGoProject/filehandling/test.txt")
	if err != nil{
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file : ", string(data))


	fileHandlingUsingBox()
}