package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println(mirroredQuery())
}

func mirroredQuery() string {
	response := make(chan string)
	cancel := make(chan struct{})
	for i := 0; i < 3; i++ {
		go request("https://www.baidu.com", cancel, response)
	}
	resp := <-response
	close(cancel)
	return resp
}

func request(hostname string, cancel <-chan struct{}, response chan<- string) {

	request, err := http.NewRequest("GET", hostname, nil)
	if err != nil {
		response <- fmt.Sprint(err)
		return
	}
	request.Cancel = cancel
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		response <- fmt.Sprint(err)
		return
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		response <- fmt.Sprint(err)
		return
	}
	response <- string(body)
}
