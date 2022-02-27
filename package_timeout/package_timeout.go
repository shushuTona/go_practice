package package_timeout

import (
	"fmt"
	"time"
)

func Executetimeout() {
	fmt.Println("===== package_timeout =====")

	chan1 := make(chan bool)
	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- true
	}()

	// select は最初に準備できた値を受信する
	// → res1が受信される
	select {
		case res1 := <- chan1:
		fmt.Println(res1)
		
		case <-time.After(time.Second * 3):
		fmt.Println("timeout chan1")
	}

	chan2 := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		chan2 <- true
	}()

	// select は最初に準備できた値を受信する
	// → res2が受信される前に、<-time.After(time.Second * 1)によってタイムアウトが起きる
	select {
		case res2 := <- chan2:
		fmt.Println(res2)
		
		case <-time.After(time.Second * 1):
		fmt.Println("timeout chan2")
	}
}
