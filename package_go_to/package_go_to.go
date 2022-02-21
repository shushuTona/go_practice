package package_go_to

import (
    "fmt"
	"time"
    "math/rand"
)

func ExecuteGoTo() int {
    rand.Seed(time.Now().UnixNano())
	rand_int := rand.Intn(10)

	fmt.Println(rand_int)

	if rand_int >5 {
		fmt.Println("rand_int >5")
	} else {
		goto ElseLabel
	}

	return rand_int

	ElseLabel:
		rand_int = 999

		return rand_int
}
