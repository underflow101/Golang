// Go Simple Timer

package simpleTimer

import (
	"fmt"
	"time"
)

func BlockingTimer(secs int) {
	for secs > 0 {
		fmt.Println(secs)
		time.Sleep(time.Second)
		secs--
	}
}

func CallBackTimer(secs int) {
	time.AfterFunc((7 * time.Second), func() {
		
	})
}

func simpleTimer() {
	fmt.Println("Start simpleTimer")
	BlockingTimer(5)
	fmt.Println("YEAE")
	CallBackTimer(7)
}
