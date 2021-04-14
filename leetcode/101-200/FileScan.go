package main

import (
	"fmt"
	"io/ioutil"
	path2 "path"
	"sync"
	"time"
)

var fileC = make(chan string, 200)
var max = 5
var gC = make(chan bool, max)
var wait = sync.WaitGroup{}

func main() {
	for i := 0; i < max; i++ {
		gC <- true
	}
	//扫描基础目录
	wait.Add(1)
	go f2("base path", true)
	wait.Wait()
}

func f2(path string, goRun bool) {
	dir, _ := ioutil.ReadDir(path)
	for _, file := range dir {
		if file.IsDir() {
			name := path2.Join(path, file.Name())
			select {
			//并发未达到最高值
			case <-gC:
				wait.Add(1)
				go f2(name, true)
			default:
				f2(name, false)
			}
		} else {
			fmt.Println(file.Name())
		}
	}
	if goRun {
		fmt.Printf("done dir %s \n", path)
		wait.Done()
		gC <- true
	}
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
