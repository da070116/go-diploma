package main

import (
	"fmt"
	"go-diploma/pkg/accidentdata"
	"go-diploma/pkg/utils"
)

func main() {
	fmt.Println("main function for Accidents data app")
	dataServer := utils.GetEnvVariable("DATASERVERPATH")
	dataServerPort := utils.GetEnvVariable("DATASERVERPORT")
	accidentService := accidentdata.GetAccidentService()
	fmt.Println(accidentService.Execute(fmt.Sprintf("http://%s:%s/accendent", dataServer, dataServerPort)))
}
