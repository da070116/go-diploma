package main

import (
	"fmt"
	"go-diploma/pkg/billingdata"
)

func main() {
	fmt.Println("main function for Conf app")

	billingService := billingdata.GetBillingService()
	fmt.Println(billingService.Execute("billing.cfg"))
}
