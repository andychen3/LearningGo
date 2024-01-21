package main

import (
	"fmt"
	"os"

	"github.com/andychen3/Files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	path := rootPath + "/data/text.txt"
	contents, err := fileutils.ReadTextFile(path)
	if err != nil {
		fmt.Println("error")
	} else {
		fmt.Println(contents)
		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", contents, contents, contents)
		fileutils.WriteToFile("output.txt", newContent)
	}
}
