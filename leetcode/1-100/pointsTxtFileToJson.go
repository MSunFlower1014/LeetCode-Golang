package main

import (
	"io/ioutil"
	"path"
	"strings"
)

func main() {
	dirPath := "D:\\Hex\\Target\\psenet\\res-image"
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return
	}
	jsonContent, err := ioutil.ReadFile(path.Join(dirPath, "000000.json"))
	if err != nil {
		return
	}

	for _, file := range dir {
		name := file.Name()
		extName := path.Ext(name)
		basePath := strings.ReplaceAll(name, extName, "")
		if extName == ".txt" {
			ioutil.WriteFile(path.Join(dirPath, basePath+".json"), jsonContent, 2)
		}
	}
}
