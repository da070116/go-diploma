package main

import (
	"fmt"
	"go-diploma/pkg/supportdata"
)

func main() {
	fmt.Println("main function for Support data app")
	supportService := supportdata.GetSupportService()

	fmt.Println(supportService.Execute("http://127.0.0.1:8383/support"))
}
