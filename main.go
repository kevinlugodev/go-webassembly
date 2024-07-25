package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("add", js.FuncOf(add))
	<-c
}

func add(this js.Value, args []js.Value) interface{}{
	sum := 0
	for _, value := range args {
		sum += value.Int()
	}
	return js.ValueOf(sum)
}