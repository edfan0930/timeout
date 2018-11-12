package main

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

const url = "http://localhost:3399"

func main() {
	SendRequest()
}

func SendRequest() (b []byte, err error) {

	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	req = req.WithContext(ctx)

	response, err := client.Do(req)
	if err != nil {
		return
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return nil, errors.New("http error")
	}
	return ioutil.ReadAll(response.Body)
}
