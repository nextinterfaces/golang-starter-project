package main

import (
	"fmt"
	"nextinterfaces/util/string_util"
	"nextinterfaces_playground/package_a"
	"nextinterfaces_playground/hello"
	"nextinterfaces_playground/main/sub_dir"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/nextinterfaces/golang-helloworld/helloworld"
)

//func (r rect) area2() float64 {
//	return r.width * r.height
//}
func main() {
	fmt.Println("Здрасти, Нацко")

	helloworld.SayHello()

	hello.Print()

	fmt.Print(bson.Marshal)


	fmt.Println(sub_dir.Public_func_four())

	fmt.Println(package_a.Public_func_one("asdasdasd"))
	fmt.Println(package_a.Public_func_two())

	fmt.Printf(string_util.Reverse("\n!oG ,olleH\n"))
	fmt.Println(string_util.Length("Computing this important string length"))

}