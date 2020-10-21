package main

func main() {
	const s = "Go101.org"
	// len(s) == 9
	// 1 << 9 == 512
	// 512 / 128 == 4

	var a byte = 1 << len(s) / 128
	println(1 << len(s[:]) / 128)

	var c byte = 1 << len(s[:])
	var b byte = 1 << len(s[:]) / 128

	println(a, b, c)

	friends := []string{"Juliet", "Emily", "Amy"}
	// The loop we generate:
	len_temp := len(friends)
	range_temp := friends // <--- 会对切片进行引用的值拷贝
	for index_temp := 0; index_temp < len_temp; index_temp++ {
		value_temp := range_temp[index_temp]
		index := index_temp
		value := value_temp
		println(index, value)
	}
}
