package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	s "strings"
)

func sendSMS(accountSID, authToken, from, to, message string) error {
	endpoint := "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"

	params := url.Values{}
	params.Set("To", to)
	params.Set("From", from)
	params.Set("Body", message)
	paramsReader := *s.NewReader(params.Encode())

	client := &http.Client{}

	req, _ := http.NewRequest("POST", endpoint, &paramsReader)
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		err = errors.New(fmt.Sprintf("http status error: %d", res.StatusCode))
		return err
	}

	return nil
}
