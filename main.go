package main

import (
	"context"
	"fmt"
	"time"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	// defer cancel()
	// wg.Add(3)
	// go subTask(ctx, &wg, "a")
	// go subTask(ctx, &wg, "b")
	// go subTask(ctx, &wg, "c")
	// wg.Wait()

	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := criticalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("critical task cancelled due to: %v\n", err)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("success", v)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := normalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("normal task cancelled due to: %v\n", err)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("success", v)
	// }()

	// wg.Wait()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(20*time.Millisecond))
	defer cancel()
	ch := deadlineTask(ctx)
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finish")
}

func deadlineTask(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible to meet deadline")
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		ch <- "hello"
	}()
	return ch
}

// func criticalTask(ctx context.Context) (string, error) {
// 	_, cancel := context.WithTimeout(context.Background(), 1200*time.Millisecond)
// 	defer cancel()
// 	t := time.NewTicker(1000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "A", nil
// }

// func normalTask(ctx context.Context) (string, error) {
// 	t := time.NewTicker(3000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "B", nil
// }

// func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
// 	defer wg.Done()
// 	t := time.NewTicker(500 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		fmt.Println(ctx.Err())
// 		return
// 	case <-t.C:
// 		t.Stop()
// 		fmt.Println(id)
// 	}
// }
