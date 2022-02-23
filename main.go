package main

import (
	"fmt"

	"github.com/shushuTona/go_practice/package_var"
	"github.com/shushuTona/go_practice/package_go_to"
	"github.com/shushuTona/go_practice/package_func"
	"github.com/shushuTona/go_practice/package_struct"
	"github.com/shushuTona/go_practice/package_interface"
)

func main() {
	// package_var
	package_var.PrintListSlice()
	package_var.PrintMap()

	// package_go_to
	package_go_to.ExecuteGoTo()

	// package_func
	package_func.ExecuteFunc(100)
	package_func.ExecuteFunc(20)
	var _, sumResult = package_func.ExecuteFunc(3)
	fmt.Println(sumResult)

	// package_struct
	register := package_struct.Register{
		RegisterId: 0,
		RegisterItemList: map[int]map[string]string{},
	}

	// package_struct
	register.SetItem(map[string]string{
		"name": "test100",
	})
	register.SetItem(map[string]string{
		"name": "test200",
	})
	registerItemList :=register.SetItem(map[string]string{
		"name": "test300",
	})
	fmt.Println(registerItemList)

	// package_interface
	person := package_interface.Person{"Tom"}
	animal := package_interface.Animal{"Dog"}
	package_interface.PrintOut(person)
	package_interface.PrintOut(animal)
}
