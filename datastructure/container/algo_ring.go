package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	//创建10个元素的闭环
	r := ring.New(10)

	//给闭环中的元素赋值
	for i := 1; i <= r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	//循环打印闭环中的元素值
	r.Do(
		func(p interface{}) {
			print(strconv.Itoa(p.(int)) + " ")
		})

	//获得当前元素之后的第5个元素
	r5 := r.Move(5)
	fmt.Println(r5)
	fmt.Println(r)

	//链接当前元素r与r5，相当于删除了r与r5之间的元素
	r1 := r.Link(r5)
	fmt.Println(r1)
	r1.Do(
		func(p interface{}) {
			print(strconv.Itoa(p.(int)) + " ")
		})
	fmt.Println(r)
}
