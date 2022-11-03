package main

import (
	"fmt"
	"go-diploma/pkg/mmsdata"
	"go-diploma/pkg/utils"
)

// main  function for MMSData app
func main() {
	fmt.Println("main function for MMS data app")
	dataServer := utils.GetEnvVariable("DATASERVERPATH")
	dataServerPort := utils.GetEnvVariable("DATASERVERPORT")
	mmsService := mmsdata.GetMMSService()
	fmt.Println(mmsService.Execute(fmt.Sprintf("http://%s:%s/mms", dataServer, dataServerPort)))
}
