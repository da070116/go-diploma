package main

import (
	"fmt"
	"go-diploma/pkg/mmsdata"
)

// main  function for MMSData app
func main() {
	fmt.Println("main function for MMS data app")
	mmsService := mmsdata.GetMMSService()
	fmt.Println(mmsService.Execute("http://127.0.0.1:8383/mms"))
}
