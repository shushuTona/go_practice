package package_channelsRange

import (
	"fmt"
)

func ExecuteChannelsRange() {
	fmt.Println("===== package_channelsRange =====")

	intChan := make(chan int, 5)

	for index := 1; index <= 5; index++ {
		intChan <- index
	}
	close( intChan )

	for intChanElem := range intChan {
		fmt.Println( intChanElem )
	}
}
