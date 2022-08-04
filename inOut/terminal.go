package inOut

import "fmt"

type terminalStruct struct {
	score int
}

func terminalInit() terminalStruct {
	return terminalStruct{0}
}

func (t terminalStruct) PutGet(str string) string {
	fmt.Println(str)
	out := ""
	fmt.Scanf("%s", &out)
	return out
}

func (t terminalStruct) Put(str string) {
	fmt.Println(str)
}
