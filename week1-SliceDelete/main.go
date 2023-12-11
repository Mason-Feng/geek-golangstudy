package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//一般方法删除int类型的切片指定元素
	sl1 := DeleteInt(sl, 1)
	fmt.Printf("sl1为: %v, sl1的长度为：%d ,sl1的容量为 %d \n", sl1, len(sl1), cap(sl1))
	sl2 := DeleteInt(sl1, 1)
	fmt.Printf("sl2为: %v, sl2的长度为：%d ,sl2的容量为 %d \n", sl2, len(sl2), cap(sl2))
	sl3 := DeleteInt(sl2, 1)
	fmt.Printf("sl3为: %v, sl3的长度为：%d ,sl3的容量为 %d \n", sl3, len(sl3), cap(sl3))
	sl4 := DeleteInt(sl3, 1)
	fmt.Printf("sl4为: %v, sl4的长度为：%d ,sl4的容量为 %d \n", sl4, len(sl4), cap(sl4))

	//使用泛型删除
	slx := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//一般方法删除int类型的切片指定元素
	slx1 := DeleteInt(slx, 1)
	fmt.Printf("slx1为: %v, slx1的长度为：%d ,slx1的容量为 %d \n", slx1, len(slx1), cap(slx1))
	slx2 := DeleteInt(slx1, 1)
	fmt.Printf("slx2为: %v, slx2的长度为：%d ,slx2的容量为 %d \n", slx2, len(slx2), cap(slx2))
	slx3 := DeleteInt(slx2, 1)
	fmt.Printf("slx3为: %v, slx3的长度为：%d ,slx3的容量为 %d \n", slx3, len(slx3), cap(slx3))
	slx4 := DeleteInt(slx3, 1)
	fmt.Printf("slx4为: %v, slx4的长度为：%d ,slx4的容量为 %d \n", slx4, len(slx4), cap(slx4))
	//使用泛型删除字符串数组
	str := []string{"f", "e", "n", "g", "q", "i", "a", "o", "w", "e", "i"}
	fmt.Printf("str:%v str的长度为：%d ,容量为 %d \n", str, len(str), cap(str))
	str1 := DeleteAny[string](str, 2)

	fmt.Printf("str1:%v str1的长度为：%d ,容量为 %d \n", str1, len(str1), cap(str1))
	str2 := DeleteAny[string](str1, 2)
	fmt.Printf("str2:%v str2的长度为：%d ,容量为 %d \n", str2, len(str2), cap(str2))
	str3 := DeleteAny[string](str2, 2)
	fmt.Printf("str3:%v str3的长度为：%d ,容量为 %d \n", str3, len(str3), cap(str3))
	str4 := DeleteAny[string](str3, 2)
	fmt.Printf("str4:%v str4的长度为：%d ,容量为 %d \n", str4, len(str4), cap(str4))
	str5 := DeleteAny[string](str4, 2)
	fmt.Printf("str5:%v str5的长度为：%d ,容量为 %d \n", str5, len(str5), cap(str5))
	str6 := DeleteAny[string](str5, 2)
	fmt.Printf("str6:%v str6的长度为：%d ,容量为 %d \n", str6, len(str6), cap(str6))

	str7 := []string{"f", "e", "n", "g", "q", "i", "a", "o", "w", "e", "i"}
	fmt.Printf("str7:%v str7的长度为：%d ,容量为 %d \n", str7, len(str7), cap(str7))
	str8, err := DeleteAt[string](str7, 2)
	if err != nil {
		return
	}
	fmt.Printf("str8:%v str8的长度为：%d ,容量为 %d \n", str8, len(str8), cap(str8))

}
