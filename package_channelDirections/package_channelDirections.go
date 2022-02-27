package package_channelDirections

import (
	"fmt"
	"time"
)

func setChanText(textChan chan<- string) {
	fmt.Println("start setChanText")

	time.Sleep(time.Second * 3)

	textChan <- "chan Text"

	fmt.Println("finish setChanText")
}

func getText(textChan <-chan string) {
	fmt.Println("start getText")

	text := <- textChan

	fmt.Println(text)

	fmt.Println("finish getText")
}

func ExecuteChannelDirections() {
	fmt.Println("===== package_channelDirections =====")

	textChan := make(chan string, 1)

	go setChanText( textChan )

	fmt.Println("setChanText...")

	getText( textChan )
}
