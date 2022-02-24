package package_defer

import (
	"fmt"
)

func DeferFunc(num int) {
    defer fmt.Println(num)

    fmt.Println("hello")
}
