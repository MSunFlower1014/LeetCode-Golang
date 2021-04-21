package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {

	path := "D:\\迅雷下载"

	dir, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range dir {
		if v.IsDir() {
			dirName := v.Name()

			if dirName == "PRED295" {
				continue
			}
			temp, err := ioutil.ReadDir(filepath.Join(path, dirName))
			if err != nil {
				fmt.Println(err)
			}
			for _, f := range temp {
				fName := f.Name()
				fPath := filepath.Join(path, dirName, fName)
				if f.Size() > 1000 {
					_ = os.Rename(fPath, filepath.Join(path, fName))
				}
			}
		}
	}
}
