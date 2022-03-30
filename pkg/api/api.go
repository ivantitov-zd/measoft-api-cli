package api

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const apiUrl = "https://home.courierexe.ru/api"

type EncodedText struct {
	XMLName xml.Name `xml:"encoding"`
	Text    string   `xml:"text"`
}

func EncodeText(text string) (string, error) {
	payload := fmt.Sprintf(
		"<?xml version=\"1.0\" encoding=\"UTF-8\"?>"+
			"<encoding><text>%s</text></encoding>",
		text,
	)
	response, err := http.Post(apiUrl, "application/xml", bytes.NewBufferString(payload))
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}

	var encodedText EncodedText
	err = xml.Unmarshal(body, &encodedText)
	if err != nil {
		log.Fatalln(err)
		return "", err
	}
	return encodedText.Text, nil
}
