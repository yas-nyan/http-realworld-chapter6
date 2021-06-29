package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

func handlerChunkedResponse(w http.ResponseWriter, r *http.Request) {
	flusher, ok := w.(http.Flusher)
	if !ok {
		panic("expected http.ResponseWriter to be an http.Flusher")
	}
	for i := 1; i <= 10; i++ {
		fmt.Fprintf(w, "Chunk #%d\n", i)
		flusher.Flush()
		time.Sleep(500 * time.Millisecond)
	}
	flusher.Flush()
}

// mainの中で、/chunkedというパスでチャンク送信を行うようにハンドラを登録

func handler(w http.ResponseWriter, r *http.Request) {
	dump, err := httputil.DumpRequest(r, true)
	if err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		return
	}
	fmt.Println(string(dump))
	fmt.Fprintf(w, "<html><body>hello</body></html>\n")
}

func main() {
	portText := ":18888"
	http.HandleFunc("/", handler)
	http.HandleFunc("/chunked", handlerChunkedResponse)
	log.Println("start http listening" + portText)
	err := http.ListenAndServe(portText, nil)
	log.Println(err)
}
