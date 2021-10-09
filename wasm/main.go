package main

import (
	"fmt"
	"syscall/js"
)

func signIn(this js.Value, args []js.Value) interface{} {
	return "Hello, " + args[0].String()
}

func readData(this js.Value, args []js.Value) interface{} {
	return js.ValueOf("Test data from WASM")
}

func registerFunctions() {
	js.Global().Set("wasmSignIn", js.FuncOf(signIn))
	js.Global().Set("wasmReadData", js.FuncOf(readData))
}

func main() {
	c := make(chan int)

	// Map Go functions to JS
	registerFunctions()

	// console.log
	fmt.Println("WASM initialized")

	// DOM manipulations
	doc := js.Global().Get("document")
	titleH1 := doc.Call("createElement", "h1")
	titleH1.Set("innerText", "WASM Demo")
	doc.Get("body").Call("appendChild", titleH1)

	// Use channel to run forever
	<-c
}
