package runtime

import (
	"testing"
)

/*
数组类型的值（以下简称数组）的长度是固定的，而切片类型的值（以下简称切片）是可变长的。

Go 语言的切片类型属于引用类型，同属引用类型的还有字典类型、通道类型、函数类型等；
而 Go 语言的数组类型则属于值类型，同属值类型的有基础数据类型以及结构体类型。

调用内建函数len，得到数组和切片的长度。通过调用内建函数cap，我们可以得到它们的容量。

切片的容量实际上代表了它的底层数组的长度

切片类型源码地址：src/runtime/slice.go

type slice struct {
	array unsafe.Pointer		//指针地址
	len   int					//长度
	cap   int					//容量
}

扩容机制：
容量<1024 - 扩容为原容量的二倍
容量>1024 - 扩容为原容量的1.25倍

使用new创建返回的是一个已清零内存数组指针
使用make创建返回的是结构体，包含数组起点指针，长度和容量

Ps:数组为值类型，切片为引用类型
*/

func TestSliceNew(t *testing.T) {
	s := new([]int)
	t.Logf("new [] int , value is %v", s)
	if s == nil {
		t.Error("new [] int , value is nil")
	}

	s2 := make([]int, 100)
	t.Logf("s length is %v", len(s2))
	t.Logf("s length is %v", len(*s))
}
