package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for {
		printPrompt()
		handleInput()
		fmt.Fprint(os.Stdout, "$ ")
	}
}

func printPrompt() {
	fmt.Fprint(os.Stdout, "$ ")
}

func handleInput() {
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return
	}

	command = command[:len(command)-1]
	switch command {
	case "exit":
		return
	default:
		fmt.Printf("%v: command not found\n", command[:len(command)-1])
	}
}
