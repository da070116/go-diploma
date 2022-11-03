package main

import (
	"fmt"
	"go-diploma/pkg/supportdata"
	"go-diploma/pkg/utils"
)

func main() {
	dataServer := utils.GetEnvVariable("DATASERVERPATH")
	dataServerPort := utils.GetEnvVariable("DATASERVERPORT")
	fmt.Println("main function for Support data app")
	supportService := supportdata.GetSupportService()

	fmt.Println(supportService.Execute(fmt.Sprintf("http://%s:%s/support", dataServer, dataServerPort)))
}
