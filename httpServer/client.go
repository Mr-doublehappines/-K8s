package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.Client{}
	resp, err := client.Get("http://127.0.0.1:7848/header")
	if err != nil {
		fmt.Println("client get err:", err)

	}
	fmt.Println(resp.Header)

	version := resp.Header.Get("VERSION")
	fmt.Println(version)

	url := resp.Request.URL
	statusCode := resp.StatusCode
	status := resp.Status
	fmt.Println(url)
	fmt.Println(statusCode)
	fmt.Println(status)

	body := resp.Body
	fmt.Printf("body:%p\n",body)
	buffer , err := ioutil.ReadAll(body)
	if err != nil{
		fmt.Println("err:",err)
	}
	fmt.Println("body is:",string(buffer))

}
