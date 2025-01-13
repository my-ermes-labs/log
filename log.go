package log

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func MyLog(bodyContent string) (string, error) {
	url := "http://192.168.64.1:3000/"

	requestBody := bytes.NewBufferString(bodyContent)

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		return "", fmt.Errorf("error while creating the request: %v", err)
	}

	// Imposta l'header Content-Type per indicare che stiamo inviando testo semplice
	req.Header.Set("Content-Type", "text/plain")

	// Invia la richiesta con un client HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error while sending the request: %v", err)
	}
	defer resp.Body.Close()

	// Legge la risposta del server
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error while reading the response: %v", err)
	}

	return string(responseBody), nil
}
