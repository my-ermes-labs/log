package log

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func MyLog(data string) {
	// URL del server Axum (modifica l'indirizzo IP della tua macchina virtuale)
	url := "http://10.168.89.115:3000/" // Modifica l'IP e il percorso in base al tuo caso

	// Corpo della richiesta (puoi inviare la stringa come JSON o come corpo di una richiesta normale)
	body := []byte(data) // Convertiamo la stringa in un array di byte

	// Creiamo una richiesta POST
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Fatal(err)
	}

	// Impostiamo l'header Content-Type
	req.Header.Set("Content-Type", "text/plain")

	// Inviamo la richiesta
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Stampiamo il codice di stato della risposta
	fmt.Println("Response Status:", resp.Status)
}
