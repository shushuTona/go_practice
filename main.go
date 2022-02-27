package main

import (
	"fmt"

	"github.com/shushuTona/go_practice/package_var"
	"github.com/shushuTona/go_practice/package_go_to"
	"github.com/shushuTona/go_practice/package_func"
	"github.com/shushuTona/go_practice/package_struct"
	"github.com/shushuTona/go_practice/package_interface"
	"github.com/shushuTona/go_practice/package_interface_type"
	"github.com/shushuTona/go_practice/package_pointer"
	"github.com/shushuTona/go_practice/package_closure"
	"github.com/shushuTona/go_practice/package_defer"
	"github.com/shushuTona/go_practice/package_goroutines"
	"github.com/shushuTona/go_practice/package_channels"
	"github.com/shushuTona/go_practice/package_channelBuffering"
	"github.com/shushuTona/go_practice/package_channelSync"
	"github.com/shushuTona/go_practice/package_channelDirections"
	"github.com/shushuTona/go_practice/package_select"
	"github.com/shushuTona/go_practice/package_timeout"
	"github.com/shushuTona/go_practice/package_closingChannels"
	"github.com/shushuTona/go_practice/package_channelsRange"
)

func main() {
	// package_var
	package_var.PrintListSlice()
	package_var.PrintMap()

	// slicing
	testSlice := []int{100, 200, 300, 400, 500}
	fmt.Println(testSlice)
	testSlice = testSlice[1:3]
	fmt.Println(testSlice)

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
	
	// package_pointer
	package_pointer.TestPointer(100)

	var num = 100
	package_pointer.ChangeDoubleArg(&num)
	fmt.Println(num)

	// package_closure
	var addNum = package_closure.ClosuresFunc(100)
	addNum(10)
	addNum(20)
	addNum(30)

	// package_defer
	package_defer.DeferFunc(100);

	// package_goroutines
	package_goroutines.ExecuteGoroutine()

	// package_channels
	package_channels.ExecuteGoroutineChannel()

	// package_channelBuffering
	package_channelBuffering.ExecuteChannelBuffering()

	// package_channelSync
	package_channelSync.ExecuteChannelSync()

	// package_channelDirections
	package_channelDirections.ExecuteChannelDirections()

	// package_select
	package_select.ExecuteSelect()

	// package_timeout
	package_timeout.Executetimeout()

	// package_closingChannels
	package_closingChannels.ExecuteClosingChannels()

	// package_channelsRange
	package_channelsRange.ExecuteChannelsRange()
}
