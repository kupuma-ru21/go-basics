package main

import (
	"fmt"
	"unsafe"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

type controller interface {
	speedUp() int
	speedDown() int
}

type vehicle struct {
	speed       int
	enginePower int
}

type bicycle struct {
	speed      int
	humanPower int
}

func main() {
	v := &vehicle{0, 5}
	speedUpAndDown(v)

	b := &bicycle{0, 5}
	speedUpAndDown(b)

	fmt.Println(v)

	var i1 interface{}
	var i2 any
	fmt.Printf("%[1]v %[1]T %v\n", i1, unsafe.Sizeof(i1))
	fmt.Printf("%[1]v %[1]T %v\n", i2, unsafe.Sizeof(i2))
	checkType(i2)
	i2 = 1
	checkType(i2)
	i2 = "hello"
	checkType(i2)
}

func (v *vehicle) speedUp() int {
	v.speed = v.speed + (10 * v.enginePower)
	return v.speed
}

func (v *vehicle) speedDown() int {
	v.speed = v.speed - (5 * v.enginePower)
	return v.speed
}

func (b *bicycle) speedUp() int {
	b.speed = b.speed + (3 * b.humanPower)
	return b.speed
}

func (b *bicycle) speedDown() int {
	b.speed = b.speed - b.humanPower
	return b.speed
}

func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}

func (v *vehicle) String() string {
	// NOTE: Sprintfは標準出力でなく、単なるstringとして出力される
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
}

func checkType(i any) {
	switch i.(type) {
	case nil:
		fmt.Println("nil")
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("unknown")
	}
}
