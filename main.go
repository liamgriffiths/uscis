package main

import (
	"flag"
	"log"

	raven "github.com/getsentry/raven-go"
)

func main() {
	var caseNumber = flag.String("caseNumber", "", "13-letter case number")
	var accountSID = flag.String("accountSID", "", "Twilio account SID")
	var authToken = flag.String("authToken", "", "Twilio auth token")
	var from = flag.String("from", "", "Twilio 'From' phone number")
	var to = flag.String("to", "", "Twilio 'To' phone number")
	var sentryDSN = flag.String("sentryDSN", "", "Log errors with Sentry.io using this DSN")
	flag.Parse()

	if *caseNumber == "" || len(*caseNumber) != 13 {
		log.Fatalln("Error: -caseNumber is required")
	}

	if *accountSID == "" {
		log.Fatalln("Error: -accountSID is required")
	}

	if *authToken == "" {
		log.Fatalln("Error: -authToken is required")
	}

	if *from == "" {
		log.Fatalln("Error: -from is required")
	}

	if *to == "" {
		log.Fatalln("Error: -to is required")
	}

	if *sentryDSN != "" {
		raven.SetDSN(*sentryDSN)
	}

	status, err := getStatus(*caseNumber)
	if err != nil {
		log.Fatal(err)
	}
	err = sendSMS(*accountSID, *authToken, *from, *to, status)
	if err != nil {
		log.Fatal(err)
	}
}
