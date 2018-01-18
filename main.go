package main

import (
	"io/ioutil"
	"os"

	"github.com/josa42/go-terminal-markdown/markdown"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		data, _ := ioutil.ReadAll(os.Stdin)
		markdown.Print(string(data))
	} else {
		for _, filePath := range args {
			data, _ := ioutil.ReadFile(filePath)
			markdown.Print(string(data))
		}
	}

}
