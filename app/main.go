package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var shellBuiltins = map[string]int{
	"exit": 0,
	"echo": 1,
	"type": 2,
}

func main() {

	for {
		printPrompt()
		exit := handleInput()
		if exit == 1 {
			break
		}
	}
}

func printPrompt() {
	fmt.Fprint(os.Stdout, "$ ")
}

func handleInput() (exit int){
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return 0
	}

	command = command[:len(command)-1]
	commandWords := strings.Split(command, " ")

	switch commandWords[0] {
	case "exit":
		return 1
	case "echo":
		fmt.Println(strings.Join(commandWords[1:], " "))
	case "type":
		if _, ok := shellBuiltins[commandWords[1]]; ok {
			fmt.Printf("%s is a shell builtin\n", commandWords[1])
		} else {
			fmt.Printf("%s: not found\n", commandWords[1])
		}
	default:
		fmt.Printf("%v: command not found\n", command)
	}
	return 0
}
