package main

import "fmt"

// 删除一个int类型切片中一个元素
func DeleteInt(slice []int, idx int) []int {
	if idx < 0 || idx >= len(slice) {
		fmt.Println("idx不合法")
	}
	var slt []int
	sltlen := len(slice) - 1
	sltcap := ReduceCap[int](slice, sltlen)

	slice = append(slice[:idx], slice[idx+1:]...)
	slt = make([]int, sltlen, sltcap)
	copy(slt, slice[:sltlen])
	return slt

}

// 缩容机制
func ReduceCap[T any](slice []T, slicelength int) int {

	slicecap := cap(slice)
	switch {
	case slicelength > (slicecap/4*3) && slicelength <= slicecap:
		return slicecap
	case slicelength > (slicecap/2) && slicelength <= (slicecap/4*3):
		return slicecap / 4 * 3
	case slicelength <= (slicecap / 2):
		return slicecap / 2
	default:
		return slicelength
	}
}

// 泛型删除
func DeleteAny[T any](slice []T, idx int) []T {
	if idx < 0 || idx >= len(slice) {
		fmt.Println("idx不合法")
	}
	var slt []T
	sltlen := len(slice) - 1
	sltcap := ReduceCap[T](slice, sltlen)

	slice = append(slice[:idx], slice[idx+1:]...)
	slt = make([]T, sltlen, sltcap)
	copy(slt, slice[:sltlen])
	return slt
}
