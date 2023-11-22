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

	d.Add("golang", "A wonderful language")
	entry, _ := d.Get("golang")
	fmt.Println(entry)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error :%v\n", err)
		os.Exit(1)
	}
}
