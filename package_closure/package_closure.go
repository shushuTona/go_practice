package package_closure

import (
	"fmt"
)

func ClosuresFunc(num int) func(int) int {
	initNum := num

	return func(addNum int) int {
		initNum = initNum + addNum
		fmt.Println(initNum)

		return initNum
	}
}