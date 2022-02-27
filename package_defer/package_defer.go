package package_defer

import (
	"fmt"
)

func DeferFunc(num int) {
    fmt.Println("===== package_defer =====")

    defer fmt.Println(num)

    fmt.Println("hello")
}
