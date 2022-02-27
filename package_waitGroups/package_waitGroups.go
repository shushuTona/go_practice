package package_waitGroups

import (
	"fmt"
	"sync"
	"time"
)

func worker( num int, wg *sync.WaitGroup ) {
	// 関数がreturnされるタイミングでwgは完了される
	defer wg.Done()

	fmt.Printf( "Worker %d starting\n", num )

	time.Sleep( time.Second )

	fmt.Printf( "Worker %d done\n", num )
}

func ExecuteWaitGroups() {
	fmt.Println("===== package_waitGroups =====")

	// WaitGroup定義
	var wg sync.WaitGroup

	for index := 1; index <= 5; index++ {
		wg.Add( 1 )

		go worker( index, &wg )
	}

	// 全てのworker処理が完了するのを待つ
	wg.Wait()
}
