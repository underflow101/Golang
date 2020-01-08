// Go Routine Control

package main

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil
}

func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}

func main() {
	runtime.GOMAXPROCS(4)
	var wait sync.WaitGroup
	var urls = []string{
		"http://xkxqjlzvieat874751.gcdn.ntruss.com/2/2019/d265/2d2651001bb575d64812b398661b39589500a9084c38a772f4b409035f74bf4e5_o_st.jpg",
		"https://file.mk.co.kr/meet/neds/2019/06/image_readtop_2019_441884_15610734753796599.jpg",
		"https://t1.daumcdn.net/news/201806/15/seouleconomy/20180615151617982vtvw.jpg",
	}

	for _, url := range urls {
		wait.Add(1)
		go func(url string) {
			defer wait.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wait.Wait()

	filenames, err := filepath.Glob("*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = writeZip("kei_img.zip", filenames)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE!")
}
