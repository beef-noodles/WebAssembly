// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
)

func getHashTest() (interface{}, error) {
	file, _ := os.Open("blob:http://127.0.0.1:8080/487bde45-d2f0-49fc-bf4c-7641d9d02576")
	// file, _ := os.Open("/Users/c4/Desktop/Personal/01.Project/web_assembly/src/go/test.go")
	defer file.Close()
	println(reflect.TypeOf(file).String())
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		fmt.Println(err)
		return nil, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	md5String := hex.EncodeToString(hashInBytes)
	fmt.Println(md5String)
	return md5String, nil
}

func main() {
	fmt.Println("Hello, 世界")
	getHashTest()
}
