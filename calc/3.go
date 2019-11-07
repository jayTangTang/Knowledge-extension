/*
@Time : 2019/11/7 11:09
@Author : ASUS
@File : 3
@Software : GoLand
*/

package main

import "fmt"

func Addd(a , b int) int {
	return  a + b
}

func Sub(a , b int) int {
	return  a - b
}

func main() {
	fmt.Println("calc.Addd called!")
	res := Addd(10, 20)
	fmt.Println("Addd(10 ,20) :", res)
	fmt.Println("calc.Sub called!")
	res = Sub(30, 20)
	fmt.Println("Sub(30 ,20) :", res)
}
