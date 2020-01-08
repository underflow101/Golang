// Go Routine Control

package main
import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"zip"
	"net/http"
	"os"
	"io"
)



func main() {
	runtime.GOMAXPROCS(4)
	var wait sync.WaitGroup
	var urls = []string{
		"http://xkxqjlzvieat874751.gcdn.ntruss.com/2/2019/d265/2d2651001bb575d64812b398661b39589500a9084c38a772f4b409035f74bf4e5_o_st.jpg"
		"https://file.mk.co.kr/meet/neds/2019/06/image_readtop_2019_441884_15610734753796599.jpg"
		"https://t1.daumcdn.net/news/201806/15/seouleconomy/20180615151617982vtvw.jpg"
	}
	
	for _, url := range urls {
		wait.Add(1);
		fmt.Println()
		go func(url string) {
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
}