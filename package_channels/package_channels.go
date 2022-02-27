package package_channels

import (
	"fmt"
)

func ExecuteGoroutineChannel() {
    fmt.Println("===== package_channels =====")

    messageChannel := make(chan string)

    go func(){
		messageChannel <- "message"
	}()

    text := <-messageChannel

    fmt.Println( text )
}
