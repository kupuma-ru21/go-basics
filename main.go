package main

import (
	"context"
	"fmt"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	numbers := []int{1, 2, 3, 4, 5}
	var i int
	flag := true
	// for v := range double(ctx, offset(ctx, double(ctx, generator(ctx, numbers...)))) {
	// 	if i == 1 {
	// 		cancel()
	// 		flag = false
	// 	}
	// 	if flag {
	// 		fmt.Println(v)
	// 	}
	// 	i++
	// }

	// for v := range generator(ctx, numbers...) {
	// 	if i == 1 {
	// 		cancel()
	// 		flag = false
	// 	}
	// 	if flag {
	// 		fmt.Println(v)
	// 	}
	// 	i++
	// }

	// for v := range double(ctx, generator(ctx, numbers...)) {
	// 	if i == 1 {
	// 		cancel()
	// 		flag = false
	// 	}
	// 	if flag {
	// 		fmt.Println(v)
	// 	}
	// 	i++
	// }

	for v := range offset(ctx, double(ctx, generator(ctx, numbers...))) {
		if i == 1 {
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

func double(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * 2:
			}
		}
	}()
	return out
}

func offset(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n + 2:
			}
		}
	}()
	return out
}
