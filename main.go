package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	// var rwMu sync.RWMutex
	// var wg sync.WaitGroup
	// var c int
	// wg.Add(4)
	// go write(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// go read(&rwMu, &wg, &c)
	// wg.Wait()
	// println("finish")

	var wg sync.WaitGroup
	var c int64
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 10; j++ {
				atomic.AddInt64(&c, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(c)
	fmt.Println("finish")
}

func read(rwMu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	time.Sleep(10 * time.Millisecond)
	rwMu.RLock()
	defer rwMu.RUnlock()
	fmt.Println("read lock")
	fmt.Println(*c)
	time.Sleep(time.Second)
	fmt.Println("read unlock")
}

func write(rwMu *sync.RWMutex, wg *sync.WaitGroup, c *int) {
	defer wg.Done()
	rwMu.Lock()
	defer rwMu.Unlock()
	fmt.Println("write lock")
	*c = *c + 1
	time.Sleep(time.Second)
	fmt.Println("write unlock")
}
