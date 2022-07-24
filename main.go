package main

import (
	"fmt"
	"os"
)

func main() {
	go preloadImages()
	bot()
}

func PrintErr(message string, err error, exit bool) {
	errorPrint := fmt.Errorf(message+": %v ", err)
	fmt.Println(errorPrint)
	if exit {
		os.Exit(2)
	}
}
