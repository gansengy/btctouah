package main

import "github.com/shopspring/decimal"

type BtcRequestBody struct {
	Price decimal.Decimal `json:"price"`
}

type EmailRequestBody struct {
	Email string `json:"email"`
}

type EmailSendStatus struct {
	Email      string `json:"email"`
	SendStatus bool   `json:"sendStatus"`
}
