package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.LstdFlags)
	ctx, cancel := context.WithTimeout(context.Background(), 600*time.Millisecond)
	defer cancel()
	const wdtTimeout = 800 * time.Millisecond
	const beatInterval = 500 * time.Millisecond
	heartBeat, v := task(ctx, beatInterval)
loop:
	for {
		select {
		case _, ok := <-heartBeat:
			if !ok {
				break loop
			}
			fmt.Println("beat pulse ⚡︎")
		case r, ok := <-v:
			if !ok {
				break loop
			}
			t := strings.Split(r.String(), "m=")
			fmt.Printf("value: %v [s]\n", t[1])
		case <-time.After(wdtTimeout):
			errorLogger.Println("doTask goroutine's heartBeat stooped")
			break loop
		}
	}
}

func task(ctx context.Context, beatInterval time.Duration) (<-chan struct{}, <-chan time.Time) {
	heartBeat := make(chan struct{})
	out := make(chan time.Time)
	go func() {
		defer close(heartBeat)
		defer close(out)
		pulse := time.NewTicker(beatInterval)
		task := time.NewTicker(beatInterval * 2)
		sendPulse := func() {
			fmt.Println("sendPulse")
			select {
			case heartBeat <- struct{}{}:
			default:
			}
		}
		sendValue := func(t time.Time) {
			for {
				select {
				case <-ctx.Done():
					return
				case <-pulse.C:
					sendPulse()
				case out <- t:
					return
				}
			}
		}
		for {
			select {
			case <-ctx.Done():
				return
			case value := <-pulse.C:
				fmt.Println("test")
				sendPulse()
				fmt.Println("value", value)
			case t := <-task.C:
				sendValue(t)
			}
		}
	}()
	return heartBeat, out
}
