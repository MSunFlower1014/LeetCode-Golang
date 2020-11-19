package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

/*
修改指定目录下xml文件内容
*/
type voc struct {
	XMLName    xml.Name     `xml:"annotation"`
	FileName   string       `xml:"filename"`
	Folder     string       `xml:"folder"`
	SizeWidth  int          `xml:"size>width"`
	SizeHeight int          `xml:"size>height"`
	SizeDepth  int          `xml:"size>depth"`
	Objects    []annotation `xml:"object"`
}

type annotation struct {
	XMLName   xml.Name `xml:"object"`
	Name      string   `xml:"name"`
	Truncated int      `xml:"truncated"`
	Difficult int      `xml:"difficult"`
	Xmin      int      `xml:"bndbox>xmin"`
	Ymin      int      `xml:"bndbox>ymin"`
	Xmax      int      `xml:"bndbox>xmax"`
	Ymax      int      `xml:"bndbox>ymax"`
}

func main() {
	files, err := ioutil.ReadDir("D:\\Hex\\Target\\psenet\\res-txt")
	if err != nil {
		return
	}
	i := 0
	for _, file := range files {
		os.Rename("D:\\Hex\\Target\\psenet\\res-txt\\"+file.Name(), "D:\\Hex\\Target\\psenet\\res-txt\\"+(string(i))+".txt")
		i++
	}
	//basePath := "D:\\Hex\\Work-Order\\印章识别\\reset\\re-sign\\sign\\Annotations"
	//files, error := ioutil.ReadDir(basePath)
	//if error != nil {
	//	fmt.Print(error)
	//}
	//
	//for _, file := range files {
	//	v := voc{}
	//	fileName := file.Name()
	//	fullPath := basePath + "/" + fileName
	//	content, error := ioutil.ReadFile(fullPath)
	//	if error != nil {
	//		fmt.Print(error)
	//		continue
	//	}
	//	err := xml.Unmarshal([]byte(content), &v)
	//	if err != nil {
	//		fmt.Printf("error: %v", err)
	//		return
	//	}
	//	fileName = strings.ReplaceAll(fileName, "xml", "jpg")
	//	v.FileName = fileName
	//
	//	output, err := xml.MarshalIndent(v, "", "\t")
	//	if err != nil {
	//		fmt.Printf("error: %v\n", err)
	//	}
	//
	//	err = ioutil.WriteFile(fullPath, output, 0666)
	//	if err != nil {
	//		fmt.Printf("error: %v\n", err)
	//	}
	//}
}
