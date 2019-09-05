package main

import (
	"syscall/js"
)

// func getFileSize(filePaths []js.Value) int64 {
// 	println(filePaths[0].String())
// 	fileInfo, _ := os.Stat(filePaths[0].String())
// 	// fileInfo, _ := os.Stat("./go/test.wasm")
// 	//文件大小
// 	filesize := fileInfo.Size()
// 	println(filesize)
// 	return filesize
// }

func add(i []js.Value) int {
	result := i[0].Int() + i[1].Int()
	js.Global().Set("output", js.ValueOf(result))
	println(js.ValueOf(result).String())
	return result
}

func subtract(i []js.Value) int {
	result := i[0].Int() - i[1].Int()
	js.Global().Set("output", js.ValueOf(result))
	println(js.ValueOf(result).String())
	return result
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// println("hello callback")
		add(args)
		return nil
	}))
	js.Global().Set("subtract", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		subtract(args)
		return nil
	}))
	// js.Global().Set("getFileSize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
	// 	getFileSize(args)
	// 	return nil
	// }))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
