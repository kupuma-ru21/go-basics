package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	eg, ctx := errgroup.WithContext(ctx)
	s := []string{"task1", "task2", "task3", "task4"}
	for _, v := range s {
		task := v
		eg.Go(func() error { return doTask(ctx, task) })
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("error %v \n", err)
	}
	fmt.Println("finish")
}

func doTask(ctx context.Context, task string) error {
	var t *time.Ticker
	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)
	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)

	}
	select {
	case <-ctx.Done():
		fmt.Printf("%v canceled : %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		fmt.Printf("task %v completed\n", task)
	}
	return nil
}

// func oldDoTask(ctx context.Context, task string) error {
// 	if task == "fake1" || task == "fake2" {
// 		return fmt.Errorf("%v failed", task)
// 	}
// 	fmt.Printf("task %v completed\n", task)
// 	return nil
// }
