package main

import (
	"fmt"
	"go-diploma/pkg/emaildata"
	"log"
)

// main function for EmailData app
func main() {
	fmt.Println("main function for EmailData app")
	emailService := emaildata.GetEmailService()
	bytes, err := emailService.ReadCSVFile("email.csv")
	if err != nil {
		log.Fatalln("no data")
	}
	err = emailService.SetData(bytes)
	if err != nil {
		return
	}
}
