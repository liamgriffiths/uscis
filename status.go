package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	s "strings"

	"github.com/PuerkitoBio/goquery"
)

func getStatus(caseNumber string) (string, error) {
	endpoint := "https://egov.uscis.gov/casestatus/mycasestatus.do"

	form := url.Values{}
	form.Set("appReceiptNum", caseNumber)

	res, err := http.PostForm(endpoint, form)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		err = errors.New(fmt.Sprintf("http status error: %d", res.StatusCode))
		return "", err
	}

	return getMessageFromResponse(res.Body)
}

func getMessageFromResponse(r io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return "", err
	}

	// hopefully this doesn't change!
	selector := "body > div.main-content-sec.pb40 > form > div > div.container > div > div > div.col-lg-12.appointment-sec.center > div.rows.text-center > p"
	text := doc.Find(selector).Text()
	if text == "" {
		err = errors.New("Couldn't find the status message in the HTML")
		return "", err
	}

	message := s.TrimSpace(text)
	return message, nil
}
