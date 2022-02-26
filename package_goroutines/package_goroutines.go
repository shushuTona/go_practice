package package_goroutines

import (
	"fmt"
    "time"
)

func printText(labelText string) {
    for index := 0; index < 3; index++ {
        fmt.Println(labelText, index)
    }
}

func ExecuteGoroutine() {
    printText("通常実行")
    go printText("goroutines1")
    go printText("goroutines2")

    time.Sleep(time.Second)
}
