package main

import (
	"fmt"
	"go-diploma/pkg/smsdata"
	"log"
)

func main() {
	fmt.Println("main function for SMSData app")
	smsService := smsdata.GetService()
	bytes, err := smsService.ReadCSVFile("sms.csv")
	if err != nil {
		log.Fatalln("no data")
	}
	err = smsService.SetData(bytes)
	if err != nil {
		return
	}
}
