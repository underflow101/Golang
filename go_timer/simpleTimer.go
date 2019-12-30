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
	time.AfterFunc
}

func SimpleTimer() {
	fmt.Println("Start simpleTimer")
	Countdown(5)
}
