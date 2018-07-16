package gFile

import (
	"fmt"
	"os"
)

//file dir auto create
func DirAutoCreate(path string) bool {
	if IsExist(path) {
		return false
	} else {
		fmt.Println("FilePathAutoCreate.noExist", path)

		err := os.MkdirAll(path, 0755)

		if err != nil {
			fmt.Println("Error", "FilePathAutoCreate", err)
			return false
		}

		return true
	}
}
