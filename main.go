package main

import (
	"fmt"

	"github.com/shushuTona/go_practice/package_var"
	"github.com/shushuTona/go_practice/package_go_to"
	"github.com/shushuTona/go_practice/package_func"
)

var executeType = false

func main() {
	if executeType {
		package_var.PrintListSlice()
		package_var.PrintMap()
		package_go_to.ExecuteGoTo()
	}

	package_func.ExecuteFunc(100)
	package_func.ExecuteFunc(20)
	var _, sumResult = package_func.ExecuteFunc(3)
	fmt.Println(sumResult)
}
