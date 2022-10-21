package main

import (
	"fmt"
	"go-diploma/pkg/billingdata"
)

func main() {
	fmt.Println("main function for EmailData app")

	billingService := billingdata.GetBillingService()
	fmt.Println(billingService.Execute("conf/billing.cfg"))
}
