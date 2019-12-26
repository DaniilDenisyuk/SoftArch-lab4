package main
 //V1 Тут можна реалізувати через конструктор, щоб CHILD автоматично ставав false...
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//ПОТРІБНО ЗМІНИТИ ВІДПОВІДНО ДО ВАШОГО GOPATH
	"github.com/DaniilDenysiuk/SoftArchLab4_v2/engine"
)

type printCommand struct {
	arg   string
	child bool
}

func (p printCommand) Execute(loop engine.Handler) {
	fmt.Println(p.arg)
}

func (p printCommand) IsChild() bool {
	return p.child
}

type addCommand struct {
	arg1, arg2 int
	child      bool
}

func (add addCommand) Execute(loop engine.Handler) {
	res := add.arg1 + add.arg2
	loop.Post(&printCommand{arg: strconv.Itoa(res), child: true})
}

func (add addCommand) IsChild() bool {
	return add.child
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
				return printCommand{"SYNTAX ERROR: " + err.Error(), false}
			}
			nums[i] = num
		}
		cmd := addCommand{nums[0], nums[1], false}
		return cmd
	}
	return printCommand{"UNKNOWN COMMAND: " + command, false}
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
		eventLoop.AwaitFinish()
		eventLoop.Post(&addCommand{2, 8, false})
	} else {
		fmt.Println(err.Error())
	}
}
