package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"reflect"
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getHash(filePath js.Value) interface{} {
	var md5String string
	var newFilePath = js.ValueOf(filePath).String()
	file, err := os.Open(newFilePath)
	println(reflect.TypeOf(file))
	checkErr(err)
	defer file.Close()
	hash := md5.New()
	if _, err = io.Copy(hash, file); err != nil {
		return err
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5String = hex.EncodeToString(hashInBytes)
	js.Global().Set("output", js.ValueOf(md5String))
	return md5String
}

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
	js.Global().Set("getHash", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		getHash(args[0])
		return nil
	}))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
