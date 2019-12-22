package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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
			if num, err := strconv.Atoi(arg); err != nil {
				return printCommand{"SYNTAX ERROR: " + err.Error()}
			} else {
				nums[i] = num
			}
			cmd := addCommand{nums[0], nums[1]}
			return cmd
		}
		return printCommand{"UNKNOWN COMMAND: " + command}
	}
	return nil
}

func main() {
	fmt.Println("in main")
	inputFile := "github.com/DaniilDenysiuk/SoftArchLab4/inputFile"
	eventLoop := new(engine.EventLoop)
	fmt.Println("in main")
	eventLoop.Start()
	if input, err := os.Open(inputFile); err == nil {
		defer input.Close()
		scanner := bufio.NewScanner(input)
		for scanner.Scan() {
			commandLine := scanner.Text()
			cmd := parse(commandLine) // parse the line to get an instance of Command
			eventLoop.Post(cmd)
		}
	}
	eventLoop.AwaitFinish()
}
