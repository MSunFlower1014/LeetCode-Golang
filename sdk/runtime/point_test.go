package runtime

import (
	"testing"
	"unsafe"
)

/*
指针转换规则
一个指针值（比如*Dog类型的值）可以被转换为一个unsafe.Pointer类型的值，反之亦然。
一个uintptr类型的值也可以被转换为一个unsafe.Pointer类型的值，反之亦然。
一个指针值无法被直接转换成一个uintptr类型的值，反过来也是如此。
*/
type User struct {
	Name string
	Age  int
}

func TestPoint(t *testing.T) {
	u := &User{Name: "ma", Age: 10}

	p := uintptr(unsafe.Pointer(u))

	offset := unsafe.Offsetof(u.Name)
	ageOffset := unsafe.Offsetof(u.Age)
	nameP := p + offset

	name := (*string)(unsafe.Pointer(nameP))
	age := (*int)(unsafe.Pointer(p + ageOffset))
	t.Logf("name : %v , age : %v\n", *name, *age)

	a := 0

	t.Logf("point : %v , value :  %v , * value : %v", &a, *&a, *&*&a)

}
