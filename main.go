package main

import (
	"dictionary/dictionary"
	"fmt"
	"os"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error :%v\n", err)
		os.Exit(1)
	}
}
