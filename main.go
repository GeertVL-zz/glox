package main

import (
	"os"
	"fmt"
	"io/ioutil"
	"bytes"
	"bufio"
)

func main() {
	argLength := len(os.Args)
	if argLength > 1 {
		fmt.Println("Usage: glox [script]")
	} else if argLength == 1 {
		runFile(os.Args[0])
	} else {
		runPrompt()
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for  {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		run(text)
	}
}

func runFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	n := bytes.IndexByte(data, 0)

	run(string(data[:n]))
}

func run(source string) {

}
