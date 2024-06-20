package main

import (
	"fmt"
	"net/url"
)

const myURL string = "http://loc.dev:8080/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("URL Demo")
	fmt.Println("url :", myURL)
	result, err := url.Parse(myURL)
	if err != nil {
		panic(err)
	}
	fmt.Println("result.scheme :", result.Scheme)
	fmt.Println("result.Host :", result.Host)
	fmt.Println("result.Path :", result.Path)
	fmt.Println("result.Raw Query :", result.RawQuery)
	fmt.Println("result.Port :", result.Port())

	query_param := result.Query()
	fmt.Println("query param :", query_param)

	fmt.Println("query_param[", "coursename", "]", query_param["coursename"])

	//construct URL
	parseOfUrl := &url.URL{
		Scheme:  "http",
		Host:    "loc.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherURL := parseOfUrl.String()
	fmt.Println("another constructed URL :", anotherURL)
}
