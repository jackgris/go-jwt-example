package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// First we need to recreate this request:
	// curl --data "user=john&pass=doe" http://localhost:3000/login

	// Make the post request:
	resp, err := http.PostForm("http://localhost:3000/login",
		url.Values{"user": {"john"}, "pass": {"doe"}})
	if err != nil {
		fmt.Println(fmt.Errorf("Error while call login url: %s", err))
	}
	// Read the body response
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while read body response: %s", err))
	}
	// Unmarshall the response like a json
	var token struct{ Token string }
	err = json.Unmarshal(body, &token)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while unmarshall body: %s", err))
	}
	// Now we have the token needed to access the restricted section.
	fmt.Println("We receive this token that will be used to access the restricted section.")
	fmt.Println(token)

	// Second we need to recreate this request:
	// curl localhost:3000/restricted -H "Authorization: Bearer eyJhbGciOiJ9.eyJleHA9.RB3arc4-OyC2W3ReWaXAt_z2Fd"
	client := &http.Client{}

	// Make the GET request:
	req, err := http.NewRequest("GET", "http://localhost:3000/restricted", nil)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while create request: %s", err))
	}
	req.Header.Add("Authorization", "Bearer "+token.Token)
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while trying get access to restricted url: %s", err))
	}
	// Read the body response
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(fmt.Errorf("Error while read body response: %s", err))
	}
	access := bytes.NewBuffer(body).String()
	// Now we show the response
	fmt.Println("We had luck and get access: ")
	fmt.Println(access)
}
