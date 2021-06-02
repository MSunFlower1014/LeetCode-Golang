package runtime

import (
	"fmt"
	"reflect"
	"testing"
)

/*
接口变量的值并不等同于这个可被称为动态值的副本。它会包含两个指针，一个指针指向动态值，一个指针指向类型信息。
除非我们只声明而不初始化，或者显式地赋给它nil，否则接口变量的值就不会为nil。

非空接口还包括申明的方法表，进行转换时，会比较结构体的方法表和接口的方法表，符合条件后进行转换
*/
type Name struct {
	Na string
}

type Na interface {
	INa()
}

func (n Name) INa() {
	fmt.Printf(" *Name : %p", &n)
}

func TestInterface(t *testing.T) {
	a := Name{
		Na: "a",
	}
	b := a
	t.Logf("a : %p , b : %p\n", &a, &b)

	var n Na
	n = a
	ty := reflect.TypeOf(n)
	t.Logf("n : %v , ty : %v\n", n, ty)
}
