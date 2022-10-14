package main

import (
	"fmt"
	"go-diploma/pkg/mmsdata"
	"log"
)

// main  function for MMSData app
func main() {
	mmsService := mmsdata.GetMMSService()
	resp, err := mmsService.SendRequest("http://127.0.0.1:8383/mms")
	if err != nil {
		log.Fatalln(err)
	}
	err = mmsService.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(mmsService.ReturnData())
}
