package main

import (
	"fmt"
	"unsafe"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

type Task struct {
	Title    string
	Estimate int
}

func main() {
	task1 := Task{Title: "Learn Golang", Estimate: 3}
	task1.Title = "Learning Golang"
	fmt.Printf("%[1]T, %+[1]v %v \n", task1, task1.Title)

	var task2 = task1
	task2.Title = "new"
	fmt.Printf("task1 is %v task2 is %v \n", task1.Title, task2.Title)

	var task1p = &Task{Title: "Learn concurrency", Estimate: 2}
	fmt.Printf("task1p: %T %+v %v\n", task1p, *task1p, unsafe.Sizeof(task1p))
	task1p.Title = "Changed"
	fmt.Printf("task1p: %+v\n", *task1p)

	var task2p *Task = task1p
	task1p.Title = "Changed by Task2"
	fmt.Printf("task1p: %+v\n", *task1p)
	fmt.Printf("task2p: %+v\n", *task2p)

	task1.extendEstimate()
	fmt.Printf("task1 value receiver: %+v\n", task1)

	task1.extendEstimatePointer()
	fmt.Printf("task1 value receiver: %+v\n", task1)
}

func (task Task) extendEstimate() {
	task.Estimate += 10
}

func (task *Task) extendEstimatePointer() {
	task.Estimate += 10
}
