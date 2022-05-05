package main

import (
	"os"
)

/*
	下划线在代码中
*/
func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("/Users/***/Desktop/text.txt")
	defer f.Close()
	for {
		n, _ := f.Read(buf)
		if n == 0 {
			break

		}
		os.Stdout.Write(buf[:n])
	}
}
