package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key") //自身のキーペアを読み込み
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // サーバ証明書のエラーを無視
			Certificates:       []tls.Certificate{cert},
		},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get("https://localhost:18443")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
