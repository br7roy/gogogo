package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	var fileAath = "C:\\Users\\Tak\\Desktop\\gofile"
	/*	fileInfo, err := ioutil.ReadDir(fileAath)
		if err != nil {
			err.Error()
		}*/

	filepath.Walk(fileAath, func(path string, file os.FileInfo, err error) error {
		if file.IsDir() && !strings.Contains(file.Name(), "@eaDir") {
			fmt.Println(path)
		}
		return nil
	})

}
