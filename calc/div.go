/*
@Time : 2019/11/7 15:15
@Author : ASUS
@File : div
@Software : GoLand
*/

package main

func Div(a ,b int) int {
	if b == 0 {
		panic("除数不能为零")
	}
	return  a / b
}
