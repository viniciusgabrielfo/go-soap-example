package main

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SayHelloResponse struct {
	XMLName xml.Name
	Body    struct {
		XMLName          xml.Name
		SayHelloResponse struct {
			XMLName       xml.Name
			HelloResponse struct {
				XMLName xml.Name
				Message string
			}
		}
	}
}

func main() {

	helloMessage, err := callSayHello("Elfo")
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(helloMessage)
}

func callSayHello(name string) (string, error) {

	log.Println("Starting call to SayHello operation...")

	sayHelloResponse := SayHelloResponse{}

	urlService := "http://www.learnwebservices.com/services/hello"

	payload := []byte(strings.TrimSpace(`
		<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/">
			<soapenv:Header/>
			<soapenv:Body>
				<SayHello xmlns="http://learnwebservices.com/services/hello">
					<HelloRequest>
						<Name>` + name + `</Name>
					</HelloRequest>
				</SayHello>
			</soapenv:Body>
		</soapenv:Envelope>
	`))

	bodyBytes, err := soapCall(urlService, "POST", payload)
	if err != nil {
		return "", err
	}

	err = xml.Unmarshal(bodyBytes, &sayHelloResponse)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error on unmarshaler body bytes to sayHelloResponse struct: %s", err.Error()))
	}

	log.Println("Call to SayHello operation was done successfully!")
	return sayHelloResponse.Body.SayHelloResponse.HelloResponse.Message, nil
}

func soapCall(webService string, action string, payload []byte) ([]byte, error) {
	log.Println("Starting call to SOAP webservice: ", webService)

	req, err := http.NewRequest(action, webService, bytes.NewReader(payload))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error on creating request: %s", err.Error()))
	}

	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error on doing request: %s", err.Error()))
	}

	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error on unmarshaler body to bytes: %s", err.Error()))
	}

	log.Println("Call to webservice was done successfully!")
	log.Println(fmt.Sprintf("body result: %s", string(bodyBytes)))
	return bodyBytes, nil
}
