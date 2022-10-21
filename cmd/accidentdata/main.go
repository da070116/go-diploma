package main

import (
	"fmt"
	"go-diploma/pkg/accidentdata"
)

func main() {
	fmt.Println("main function for Accidents data app")
	accidentService := accidentdata.GetAccidentService()
	fmt.Println(accidentService.Execute("http://127.0.0.1:8383/accendent"))
}
