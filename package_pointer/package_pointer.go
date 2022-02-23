package package_pointer

import (
	"fmt"
)

var testInt int
var testIntPointer *int

func TestPointer(num int) {
	// &testInt : testIntのポインターの参照 = 参照渡し
	testIntPointer = &testInt

	// *testIntPointer : ポインターの中身
	*testIntPointer = num

	fmt.Println(testInt)
	fmt.Println(testIntPointer)
}

func ChangeDoubleArg(num *int) {
	*num = *num * 2
}