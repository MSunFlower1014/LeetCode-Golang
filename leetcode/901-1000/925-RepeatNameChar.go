package main

/*
925. 长按键入
你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。

你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。

"xnhtq"
"xhhttqq"

"alex"
"alexxr"

"ruooeffnvjwtsidzwkwxinisxzthwzjynzzvreapsynsqfqzcr"
"rruoooeefffnvjjwtssiiddzzwkkwwxiinniissxxztthwzjyynnzzzzvrreaapsynnsqqfqzzcrr"
*/
func main() {
	name, typed := "ruooeffnvjwtsidzwkwxinisxzthwzjynzzvreapsynsqfqzcr", "rruoooeefffnvjjwtssiiddzzwkkwwxiinniissxxztthwzjyynnzzzzvrreaapsynnsqqfqzzcrr"
	flag := isLongPressedName(name, typed)
	println(flag)
}

func isLongPressedName(name string, typed string) bool {
	m := make(map[int]int)
	nameSum := 0
	typeSum := 0
	var old int32
	var sum int
	l := make([]int32, 0)
	for _, v := range typed {
		if v == old {
			sum++
		} else {
			nameSum++
			m[nameSum] = sum
			sum = 1
			old = v
			l = append(l, old)
		}
	}
	var newValue int32
	var newSum int
	for _, c := range name {
		if c == newValue {
			newSum++
		} else {
			typeSum++
			s, ok := m[typeSum]
			if !ok {
				return false
			}
			if s < newSum {
				return false
			}
			newValue = c
			newSum = 1
			ch := l[0]
			if c != ch {
				return false
			}
			l = l[1:]
		}
	}
	//对比字母出现个数
	if nameSum != typeSum {
		return false
	}
	return true
}
