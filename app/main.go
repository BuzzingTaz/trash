package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	_exitBuiltin = iota
	_echoBuiltin
	_typeBuiltin
)

var shellBuiltins = map[string]int{
	"exit": _exitBuiltin,
	"echo": _echoBuiltin,
	"type": _typeBuiltin,
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

func handleInput() (exit int) {
	exit = 0
	command, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return exit
	}

	command = command[:len(command)-1]
	commandWords := strings.Split(command, " ")

	shellBuiltin, isShellBuiltin := shellBuiltins[commandWords[0]]
	if isShellBuiltin {
		switch shellBuiltin {
		case _exitBuiltin:
			exit = 1

		case _echoBuiltin:
			fmt.Println(strings.Join(commandWords[1:], " "))

		case _typeBuiltin:
			if _, ok := shellBuiltins[commandWords[1]]; ok {
				fmt.Printf("%s is a shell builtin\n", commandWords[1])
			} else if path, err := findExecutableInPath(commandWords[1]); err == nil {
				fmt.Printf("%s is %s\n", commandWords[1], path)
			} else {
				fmt.Printf("%s: not found\n", commandWords[1])
			}
		}
		return exit
	}

	_, err = findExecutableInPath(commandWords[0])
	if err != nil {
		fmt.Printf("%v: command not found\n", command)
		return exit
	}

	cmd := exec.Command(commandWords[0], commandWords[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	return exit
}

func findExecutableInPath(cmd string) (string, error) {
	return exec.LookPath(cmd)
}
