package package_select

import (
	"fmt"
	"time"
)

func setChanText(textChan chan<- string, text string) {
	time.Sleep(time.Second * 3)

	textChan <- text
}

func ExecuteSelect() {
	fmt.Println("===== package_select =====")

	chan1 := make(chan string, 1)
	chan2 := make(chan string, 1)

	go setChanText( chan1, "chan1" )
	go setChanText( chan2, "chan2" )

	for index := 0; index < 2; {
		select {
			case text1 := <- chan1:
				fmt.Println( text1 )
				index++
			case text2 := <- chan2:
				fmt.Println( text2 )
				index++
		}
	}
}
