package package_channelBuffering

import (
	"fmt"
)

func ExecuteChannelBuffering() {
    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
