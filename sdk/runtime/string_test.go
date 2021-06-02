package runtime

import (
	"strings"
	"testing"
)

/**
runtime/string.go
type stringStruct struct {
	str unsafe.Pointer
	len int
}

string 通过持有指针和长度获取对应数据，指针可视为byte数组中起点数据指针
但string中数据不可修改，因此特性，string和byte数组相互转换时为全量复制，而非共享数据

在Go中用rune值来表示。 内置rune类型为内置int32类型的一个别名。
在UTF-8编码中，一个码点值可能由1到4个字节组成。
比如，每个英语码点值（均对应一个英语字符）均由一个字节组成，
而每个中文字符（均对应一个中文字符）均由三个字节组成。
*/
func TestStringBuilder(t *testing.T) {
	/*
		type Builder struct {
			addr *Builder // of receiver, to detect copies by value
			buf  []byte
		}
		底层通过byte数组实现
	*/
	b := strings.Builder{}

	b.WriteString("你好")

	/*
		func (b *Builder) String() string {
			return *(*string)(unsafe.Pointer(&b.buf))
		}
	*/
	t.Logf("builder to string is %v\n", b.String())

	s := "nihao我是***"
	for _, r := range s {
		//本质上rune为int32的一个别名,转为字符串可自动识别多个字节的中文
		t.Logf("int32 : %v , string : %v\n", r, string(r))
	}
}
