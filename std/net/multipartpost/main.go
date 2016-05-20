package main

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
)

func newMultipartRequest(url string, params map[string]string) (*http.Request, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	writer.Close()
	return http.NewRequest("POST", url, body)
}

func main() {
	extraParams := map[string]string{
		"title":       "My Document",
		"author":      "zieckey",
		"description": "A document with all the Go programming language secrets",
	}
	request, err := newMultipartRequest("http://10.16.28.17:8091/echo", extraParams)
	if err != nil {
		log.Fatal(err)
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Header)
		fmt.Println(body)
	}
}
