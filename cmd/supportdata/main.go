package main

import (
	"fmt"
	"go-diploma/pkg/supportdata"
	"log"
)

func main() {
	fmt.Println("main function for Support data app")
	supportService := supportdata.GetSupportService()
	resp, err := supportService.SendRequest("http://127.0.0.1:8383/support")
	if err != nil {
		log.Fatalln(err)
	}
	err = supportService.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(supportService.ReturnData())
}
