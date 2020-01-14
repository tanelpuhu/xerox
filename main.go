package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/atotto/clipboard"
)

func main() {
	if !termutil.Isatty(os.Stdin.Fd()) {
		input, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		if err := clipboard.WriteAll(string(input)); err != nil {
			panic(err)
		}
	} else {
		text, err := clipboard.ReadAll()
		if err != nil {
			panic(err)
		}
		fmt.Print(text)
	}
}
