package main

import (
	"fmt"
	"go-diploma/pkg/smsdata"
)

// main function for SMSData app
func main() {
	fmt.Println("main function for SMSData app")
	smsService := smsdata.GetSMSService()
	smsService.Execute("sms.csv")
	fmt.Println(smsService.DisplayData())
}
