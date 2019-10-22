package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}
	//b是数组，是a的一份拷贝
	b := a
	//c是切片，是引用类型，底层数组是a
	c := a[2:5]
	for i := 0; i < len(a); i++ {
		a[i] = a[i] + 1
	}
	//改变a的值后，b是a的拷贝，b不变，c是引用，c的值改变
	fmt.Println(a) //[2,3,4]
	fmt.Println(b) //[1 2 3]
	fmt.Println(c) //[2,3,4]

}
