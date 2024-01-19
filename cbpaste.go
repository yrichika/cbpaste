package main

import (
	"fmt"
	"os"

	"golang.design/x/clipboard"
)

func main() {

	clitInitErr := clipboard.Init()
	if clitInitErr != nil {
		fmt.Println("Error initializing clipboard", clitInitErr)
		os.Exit(1)
	}

	clipContent := clipboard.Read(clipboard.FmtText)
	fmt.Print(string(clipContent))
}
