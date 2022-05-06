package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func testArray(arr [2]int) {
	fmt.Printf("arr1 : %p , %v\n", &arr, arr)
}

func testArray1(arr *[2]int) {
	fmt.Printf("arr1 : %p , %v\n", arr, *arr)
}

func array() [1024]int {
	var x [1024]int
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func slice1() []int {
	x := make([]int, 1024)
	for i := 0; i < len(x); i++ {
		x[i] = i
	}
	return x
}

func BenchmarkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		array()
	}
}

func BenchmarkSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1()
	}
}

func main() {
	var arr1 = [2]int{100, 200}
	var arr2 [2]int
	arr2 = arr1
	fmt.Printf("arr1 : %p , %v\n", &arr1, arr1)
	fmt.Printf("arr2 : %p , %v\n", &arr2, arr2)
	testArray(arr1)   // 以上两种情况内存指针都不相同
	testArray1(&arr1) // 传入指针切片
	testArray2()
}

// 切片的内部结构定义
type slice struct {
	array unsafe.Pointer // 数组指针
	len   int            // 当前切片的长度
	cap   int            // 当前切片的总量
}

func testArray2() {
	// 如果想从切片中得到一块内存地址则可以
	var arr5 = make([]int, 200)
	var point1 = &arr5[0]
	fmt.Println(point1)
	// 从go的内存当中构建一个切片
	/**
	var ptr unsafe.Pointer
	var s1 = struct {
		addr uintptr
		len  int
		cap  int
	}{uintptr(ptr), 10, 20}
	s := *(*[]byte)(unsafe.Pointer(&s1))
	*/
}
