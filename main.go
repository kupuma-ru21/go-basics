package main

import (
	"fmt"
	"unsafe"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	var ui1 uint16
	fmt.Printf("memory top address of ui1 is %p\n", &ui1)

	var ui2 uint16
	fmt.Printf("memory top address of ui2 is %p\n", &ui2)

	var p1 *uint16
	fmt.Printf("value of p1 is %v\n", p1)
	p1 = &ui1
	// value of p1 is equal to memory top address of ui1
	fmt.Printf("value of p1 is %v\n", p1)
	fmt.Printf("size of p1 is %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory top address of p1 is %p\n", &p1)
	fmt.Printf("value of ui1(dereference) is %v\n", *p1)
	*p1 = 1
	fmt.Printf("value of ui1 is %v\n", ui1)

	var pp1 **uint16 = &p1
	fmt.Printf("value of pp1 is %v\n", pp1)
	fmt.Printf("size of pp1 is %d[bytes]\n", unsafe.Sizeof(pp1))
	fmt.Printf("memory top address of pp1 is %p\n", &pp1)
	fmt.Printf("value of p1(dereference) is %v\n", *pp1)
	fmt.Printf("value of ui1(dereference) is %v\n", **pp1)

	**pp1 = 10
	fmt.Printf("value of ui1 is %v\n", ui1)
}
