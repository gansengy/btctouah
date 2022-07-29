package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/mail"
	"net/smtp"
)

func ParseEmail(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

func IsEmailValid(data []EmailRequestBody, email EmailRequestBody) (string, bool) {

	addr, status := ParseEmail(email.Email)
	if !status {
		return addr, false
	}

	for _, element := range data {
		if addr == element.Email {
			return addr, false
		}
	}
	return addr, true

}

func WriteEmailToFile(data []EmailRequestBody, email string, fileName string) {
	convertedEmail := EmailRequestBody{Email: email}

	data = append(data, convertedEmail)

	dataBytes, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(fileName, dataBytes, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func MakeEmail(email string) bool {
	from := "bitcoinsender41@gmail.com"
	password := "upduisktbsoiikna"

	toEmailAddress := email
	to := []string{toEmailAddress}

	host := "smtp.gmail.com"
	port := "587"
	address := host + ":" + port

	message := []byte("From: bitcoinsender41@gmail.com\r\n" +
		"To:" + email + "\r\n" +
		"Subject: Bitcoin Price UAH\r\n\r\n" +
		"Now bitcoin price in UAH is " + btcToUah().String() + "\r\n")

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(address, auth, from, to, message)
	if err != nil {
		return false
	}
	return true
}
