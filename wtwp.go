package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	accountSid := ""
	authToken := ""
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	v := url.Values{}
	v.Set("To", "")
	v.Set("From", "+")
	v.Set("Body", "Password is...")
	rb := *strings.NewReader(v.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Send request to Twilio API
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)

}
