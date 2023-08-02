package main

import (
	"fmt"
	"sync"
	"time"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

const bufSize = 3

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, bufSize)
	wg.Add(1)

	go func() {
		defer wg.Done()
		for i := 0; i < bufSize; i++ {
			time.Sleep(1000 * time.Millisecond)
			ch <- "hello"
		}
	}()
	for i := 0; i < 3; i++ {
		select {
		case m := <-ch:
			fmt.Println(m)
		default:
			fmt.Println("no msg arrived")
		}
		time.Sleep(1500 * time.Millisecond)
	}
}
