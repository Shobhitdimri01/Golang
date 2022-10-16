package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

)

type logWriter struct{}

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	if resp != nil {
		fmt.Println("Successfully connected")
	}

	// m := make([]byte, 10000)
	// resp.Body.Read((m))
	// fmt.Println(string(m))
	lw := logWriter{}
	io.Copy(lw,resp.Body)

	b, err := io.ReadAll(resp.Body)
	bCode := []byte(b)

	

	str := make([]string, 0)
	str = append(str, "Hello", "Bye", "go", "run")
	fmt.Println("Printing", len(str))

	err = ioutil.WriteFile("Google_resp.txt", bCode, 0666)
	if err != nil {
		fmt.Println(err.Error())
	}

}

func (logWriter) Write(bs [] byte)(int,error){
	fmt.Println(string(bs))
	var n int = len(bs)
	fmt.Println("The length of byte is : ",n)
	return n,nil
}
