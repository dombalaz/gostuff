package main

import (
	"fmt"
	"os"
)

func fileLen(in string) (int64, error) {
	f, err := os.Open(in)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		s, err := fileLen(os.Args[i])
		fmt.Printf("path: %v, size: %v, err:%v\n", os.Args[i], s, err)
	}
}
