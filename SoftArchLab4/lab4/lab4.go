package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//ПОТРІБНО ЗМІНИТИ ВІДПОВІДНО ДО ВАШОГО GOPATH
	"github.com/DaniilDenysiuk/SoftArchLab4/engine"
)

type printCommand struct {
	arg string
}

func (p printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}

type addCommand struct {
	arg1, arg2 int
}

func (add addCommand) Execute(loop engine.Handler) {
	res := add.arg1 + add.arg2
	loop.Post(&printCommand{arg: strconv.Itoa(res)})
}

func parse(commandLine string) engine.Command {
	parts := strings.Fields(commandLine)
	command := parts[0]
	args := parts[1:]
	if command == "add" {
		nums := make([]int, len(args))
		for i, arg := range args {
			num, err := strconv.Atoi(arg)
			if err != nil {
				return printCommand{"SYNTAX ERROR: " + err.Error()}
			}
			nums[i] = num
		}
		cmd := addCommand{nums[0], nums[1]}
		return cmd
	}
	return printCommand{"UNKNOWN COMMAND: " + command}
}

func main() {
	inputFile := "./inputFile"
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	} else {
		fmt.Println(err.Error())
	}
	eventLoop.AwaitFinish()
}
