package main

import (
	"fmt"
	"go-diploma/pkg/accidentdata"
	"log"
)

func main() {
	fmt.Println("main function for Accidents data app")
	accidentService := accidentdata.GetAccidentService()
	resp, err := accidentService.SendRequest("http://127.0.0.1:8383/accendent")
	if err != nil {
		log.Fatalln(err)
	}
	err = accidentService.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(accidentService.ReturnData())
}
