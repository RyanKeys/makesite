package userutils

import (
	"io/ioutil"
	"log"
	"strings"
)


func ReadFile(filename string) string {
	fileContents, err := ioutil.ReadFile(filename)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	return string(fileContents)
}

//Takes in a your desired filename and the data to put in the file as bytes.
func WriteFile(filename string, data []byte) {
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		panic(err)
	}
}

//An extension of 'WriteFile(filename, data)' to handle strings of data instead of bytes.
func WriteFileFromString(filename string, data string) {
	bytesToWrite := []byte(data)
	WriteFile(filename, bytesToWrite)
}

func GetTxtFilesFromDir(dirname string) []string {
	var txtfiles [] string 
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		if strings.Contains(f.Name(),".txt") {
			txtfiles = append(txtfiles, f.Name())
		}
	}
	return txtfiles
}

