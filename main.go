package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	cores := runtime.NumCPU()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	outChs := make([]<-chan string, cores)
	inData := generator(ctx, numbers...)
	for i := 0; i < cores; i++ {
		outChs[i] = fanOut(ctx, inData, i+1)
	}
	var i int
	flag := true
	for v := range fanIn(ctx, outChs...) {
		if i == 3 {
			cancel()
			flag = false
		}
		if flag {
			fmt.Println(v)
		}
		i++
	}
	fmt.Println("finish")
}

func generator(ctx context.Context, numbers ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range numbers {
			select {
			case <-ctx.Done():
				return
			case out <- num:
			}
		}
	}()
	return out
}

func fanOut(ctx context.Context, in <-chan int, id int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		heavyFunc := func(i int, id int) string {
			time.Sleep(200 * time.Millisecond)
			return fmt.Sprintf("result:%v (id: %v)", i*i, id)
		}
		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- heavyFunc(v, id):
			}
		}
	}()
	return out
}

func fanIn(ctx context.Context, chs ...<-chan string) <-chan string {
	var wg sync.WaitGroup
	out := make(chan string)
	multiplex := func(ch <-chan string) {
		defer wg.Done()
		for text := range ch {
			select {
			case <-ctx.Done():
				return
			case out <- text:
			}
		}
	}
	wg.Add(len(chs))
	for _, ch := range chs {
		go multiplex(ch)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
