package main

import (
	"fmt"
	"math"
	"time"
	"unsafe"
)

func main() {
	now := time.Now()
	result := GetHashCode(&now)
	fmt.Printf("result is %v\n ", result)
	index := GetIndex(result, 256)
	fmt.Printf("index is %v \n", index)
	i := Float64bits(1.234)
	fmt.Printf("uint64 is %v \n", i)
	f := math.Float32frombits(uint32(i))
	fmt.Printf("float64 is %v \n", f)
}

func GetHashCode(o interface{}) int {
	p := unsafe.Pointer(&o)
	i := uintptr(p)
	return int(i)
}

func Float64bits(f float64) uint64 {
	return *(*uint64)(unsafe.Pointer(&f))
}

func GetIndex(hash, cap int) int {
	return (hash ^ (hash >> 16)) & (cap - 1)
}
