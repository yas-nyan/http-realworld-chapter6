package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	tr := &http.Transport{
		//DisableKeepAlives: true, // 明示的にこれを書くと conection: close属性が入る． TCP セッションの使いまわしを行わなくなる．
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:18443")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // 終わったら明示的にcloseする．

	body, err := ioutil.ReadAll(resp.Body)
	//ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
