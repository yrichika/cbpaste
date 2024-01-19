package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.design/x/clipboard"
)

func main() {

	clipErr := clipboard.Init()
	if clipErr != nil {
		fmt.Println("Error initializing clipboard", clipErr)
		os.Exit(1)
	}
	// TODO:
}
