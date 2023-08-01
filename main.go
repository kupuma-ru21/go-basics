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

func main() {
	// without buffer
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()

	ch1 <- 10
	close(ch1)

	v1, ok1 := <-ch1
	fmt.Printf("%v %v\n", v1, ok1)
	wg.Wait()

	// with buffer
	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)

	v2, ok2 := <-ch2
	fmt.Printf("%v %v\n", v2, ok2)

	v3, ok3 := <-ch2
	fmt.Printf("%v %v\n", v3, ok3)

	v4, ok4 := <-ch2
	fmt.Printf("%v %v\n", v4, ok4)

	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
		time.Sleep(2 * time.Second)
	}

	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v started\n", i)
			<-nCh
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh)
	fmt.Println("unlocked by manual close")
	wg.Wait()
	fmt.Println("finish")
}

func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Println("write")
		}
	}()
	return ch
}
