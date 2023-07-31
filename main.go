package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// static checkをmanualで設定した。
// https://stackoverflow.com/questions/71101439/how-can-i-configure-the-staticcheck-linter-in-visual-studio-code

// vscode goのextensionが正常に作動しなかった
// https://formulae.brew.sh/formula/goplsをinstallする必要があった

func main() {
	funcDefer()

	files := []string{"file1.csv", "file2.csv", "file3.csv"}
	fmt.Println(trimExtension(files...))

	fileName, err := fileChecker("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fileName)

	i := 1
	func(i int) {
		fmt.Println(i)
	}(i)

	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(1))

	f2 := func(fileName string) string {
		return fileName + ".csv"
	}
	addExtension(f2, "file1")

	f3 := multiply()
	fmt.Println(f3(2))

	f4 := countUp()
	// for i := 1; i <= 5; i++ {
	// 	fmt.Printf("%v\n", f4(2))
	// }
	fmt.Printf("%v\n", f4(2))
}

func funcDefer() {
	defer fmt.Println("main func final finish")
	defer fmt.Println("main func semi finish")
	fmt.Println("Hello World")
}

func trimExtension(files ...string) []string {
	out := make([]string, 0, len(files))
	for _, file := range files {
		out = append(out, strings.TrimSuffix(file, ".csv"))
	}
	return out
}

func fileChecker(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", errors.New("file not found")
	}
	defer file.Close()
	return fileName, nil
}

func addExtension(f func(fileNameGiven string) string, fileName string) {
	fmt.Println(f(fileName))
}

func multiply() func(int) int {
	return func(i int) int {
		return i * 1000
	}
}

func countUp() func(int) int {
	count := 0
	return func(i int) int {
		count += i
		return count
	}
}
