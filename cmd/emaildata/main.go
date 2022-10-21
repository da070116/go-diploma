package main

import (
	"fmt"
	"go-diploma/pkg/emaildata"
)

// main function for EmailData app
func main() {
	fmt.Println("main function for EmailData app")
	emailService := emaildata.GetEmailService()
	fmt.Println(emailService.Execute("email.csv"))
}
