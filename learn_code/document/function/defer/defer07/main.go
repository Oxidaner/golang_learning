package main

import (
	"fmt"
	"io"
	"os"
)

func do() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}

func main() {
	do()
}
