package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func fib() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func printConetent(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		str := scanner.Text()
		fmt.Println(str)
	}
}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000000000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func main() {
	a := fib()
	printConetent(a)
}
