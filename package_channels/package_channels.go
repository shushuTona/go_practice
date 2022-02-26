package package_channels

import (
	"fmt"
)

func ExecuteGoroutineChannel() {
    messageChannel := make(chan string)

    go func(){
		messageChannel <- "message"
	}()

    text := <-messageChannel

    fmt.Println( text )
}
