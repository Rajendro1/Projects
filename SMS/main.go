package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func main() {
    url := "https://api.chat-api.com/instance<instance>/message?token=<token>"

    requestBody, err := json.Marshal(map[string]string{
        "phone":  "<whatsapp_number>",
        "body":   "<message_body>",
    })

    if err != nil {
        panic(err)
    }

    request, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
    request.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    response, err := client.Do(request)

    if err != nil {
        panic(err)
    }

    defer response.Body.Close()
    responseBody, err := ioutil.ReadAll(response.Body)

    if err != nil {
        panic(err)
    }

    fmt.Println(string(responseBody))
}
