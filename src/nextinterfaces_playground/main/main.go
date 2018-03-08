package main

import (
	"fmt"
	"nextinterfaces/util/string_util"
	"nextinterfaces_playground/package_a"
	"nextinterfaces_playground/hello"
)

//func (r rect) area2() float64 {
//	return r.width * r.height
//}
func main() {
	fmt.Println("Здрасти, Нацко")

	hello.Print()

	fmt.Println(package_a.Public_func_one("asdasdasd"))
	fmt.Println(package_a.Public_func_two())

	fmt.Printf(string_util.Reverse("\n!oG ,olleH\n"))
	fmt.Println(string_util.Length("Computing this important string length"))

}
