package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://loc.dev"

func main() {
	fmt.Println("hello LCO web request")
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	//fmt.Println("response :", res)
	defer res.Body.Close()
	databyte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(databyte)
	fmt.Println("context :", content)
}
