package main

import (
	"fmt"
	"io/ioutil"
	path2 "path"
	"sync"
	"time"
)

var fileC = make(chan string, 200)
var gC = make(chan bool, 5)
var wait = sync.WaitGroup{}

func main() {
	//扫描基础目录
	fileHandle("D:\\Hex\\Work-Order\\深度学习训练平台\\打标结算")
	wait.Wait()
}
func f() {
	for {
		select {
		case name := <-fileC:
			fileHandle(name)
		case <-time.After(5 * time.Second):
			wait.Done()
			gC <- true
			return
		}
	}
}

func fileHandle(path string) {
	select {
	//并发未达到最高值
	case <-gC:
		wait.Add(1)
		go f()
	default:
	}
	dir, _ := ioutil.ReadDir(path)
	for _, file := range dir {
		if file.IsDir() {
			name := path2.Join(path, file.Name())
			fileC <- name
		} else {
			fmt.Println(file.Name())
		}
	}
}
