package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func getPrice(context *gin.Context) {
	var btc = BtcRequestBody{
		btcToUah(),
	}
	context.IndentedJSON(http.StatusOK, btc)
}

func addEmail(context *gin.Context) {
	var newEmail EmailRequestBody

	if err := context.BindJSON(&newEmail); err != nil {
		return
	}

	data := CreateSlice("emails.json")

	addr, status := IsEmailValid(data, newEmail)

	if status {
		WriteEmailToFile(data, addr, "emails.json")
		context.IndentedJSON(http.StatusOK, newEmail)
	} else {
		context.IndentedJSON(http.StatusConflict, newEmail)
	}
}

func sendEmails(context *gin.Context) {
	var data []EmailSendStatus
	flag := false

	dataEmails := CreateSlice("emails.json")

	for _, element := range dataEmails {
		if MakeEmail(element.Email) {
			flag = true
			request := EmailSendStatus{Email: element.Email, SendStatus: true}
			data = append(data, request)
		} else {
			request := EmailSendStatus{Email: element.Email, SendStatus: false}
			data = append(data, request)
		}
	}

	if flag {
		context.IndentedJSON(http.StatusOK, data)
	} else {
		context.IndentedJSON(http.StatusConflict, data)
	}
}

func main() {
	err := CheckForFile("emails.json")
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/rate", getPrice)
	router.POST("/subscribe", addEmail)
	router.POST("/sendEmails", sendEmails)

	err = router.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
