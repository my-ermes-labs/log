package log

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func MyLog(node string, data string) {
	url := "http://192.168.64.1:3000/"

	body := []byte(node + ": " + data)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Println("Response Status:", resp.Status)
}
