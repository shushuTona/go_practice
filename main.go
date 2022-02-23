package main

import (
	"fmt"

	"github.com/shushuTona/go_practice/package_var"
	"github.com/shushuTona/go_practice/package_go_to"
	"github.com/shushuTona/go_practice/package_func"
	"github.com/shushuTona/go_practice/package_struct"
	"github.com/shushuTona/go_practice/package_interface"
	"github.com/shushuTona/go_practice/package_interface_type"
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

	// package_interface_type
	castVar := package_interface_type.CastToString( []string{} )
	fmt.Println(castVar)
	typeCheck := package_interface_type.TypeCheck( true )
	fmt.Println(typeCheck)

	package_interface_type.AddListItem( "key_1", 100 )
	package_interface_type.AddListItem( "key_2", "test_value" )
	list := package_interface_type.AddListItem( "key_3", true )
	fmt.Println(list)
}
