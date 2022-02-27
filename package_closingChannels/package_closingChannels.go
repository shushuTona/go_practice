package package_closingChannels

import (
	"fmt"
)

func ExecuteClosingChannels() {
	fmt.Println("===== package_closingChannels =====")

	jobsChan := make(chan int, 5)
	doneChan := make(chan bool)

	go func() {
		for {
			job, more := <-jobsChan

			if more {
				fmt.Println("job", job)
			} else {
				fmt.Println("finish job")
				doneChan <- true
				return
			}
		}
	}()

	for index := 1; index <= 3; index++ {
		jobsChan <- index
		fmt.Println("send job", index)
	}
	
	close( jobsChan )
	fmt.Println("send all  job")

	<- doneChan
}
