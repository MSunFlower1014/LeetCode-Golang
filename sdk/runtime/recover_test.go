package runtime

import "testing"

/**
在下面的情况下，recover函数调用的返回值为nil：
1.传递给相应panic函数调用的实参为nil；
2.当前协程并没有处于恐慌状态；
3.recover函数并未直接在一个延迟函数调用中调用。
*/

func TestRecoverNil(t *testing.T) {
	defer func() {
		a := recover()
		t.Logf("recover is %v", a)
	}()

	panic(nil)
}
