package package_channelSync

import (
	"fmt"
	"time"
)

func worker(flag chan bool) {
	fmt.Println("worker start")

	time.Sleep(time.Second * 3)

	fmt.Println("worker finish")

	flag <- true
}

func ExecuteChannelSync() {
	fmt.Println("===== package_channelSync =====")

	flag := make(chan bool, 1)

	go worker( flag )

	<-flag
}
