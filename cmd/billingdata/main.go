package main

import (
	"fmt"
	"go-diploma/pkg/billingdata"
	"log"
)

func main() {
	fmt.Println("main function for EmailData app")

	billingService := billingdata.GetBillingService()
	bytes, err := billingService.ReadFile("conf/billing.cfg")
	if err != nil {
		log.Fatalln("no data")
	}

	err = billingService.SetData(bytes)
	if err != nil {
		return
	}
	fmt.Println(billingService.ReturnData())
}
