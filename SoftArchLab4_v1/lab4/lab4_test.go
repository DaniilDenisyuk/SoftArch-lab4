package main

import (
	"strconv"
	"testing"

	//ПОТРІБНО ЗМІНИТИ ВІДПОВІДНО ДО ВАШОГО GOPATH
	"github.com/DaniilDenysiuk/SoftArchLab4_v1/engine"
)

var testArr []string

type testPrint struct {
	arg string
}

func (p testPrint) Execute(loop engine.Handler) {
	testArr = append(testArr, p.arg)
}

type testAdd struct {
	arg1, arg2 int
}

func (add testAdd) Execute(loop engine.Handler) {
	res := add.arg1 + add.arg2
	loop.Post(&testPrint{arg: strconv.Itoa(res)})
}

func TestThousandsCommands(t *testing.T) {
	eventLoop := new(engine.EventLoop)
	eventLoop.Start()
	var array [10000]testAdd
	for i := 0; i < cap(array); i++ {
		array[i].arg1 = 1
		array[i].arg2 = i
		eventLoop.Post(array[i])
	}
	eventLoop.AwaitFinish()

	i := 0
	for ; i < len(array); i++ {
		data := array[i]
		for res := range testArr {
			if (data.arg2 + 1) == res {
				break
			}
		}
	}
	if i != len(array) {
		t.Error("Error", i, len(array))
	}
}
