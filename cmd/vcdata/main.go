package main

import (
	"fmt"

	"go-diploma/pkg/vcdata"
	"log"
)

// main function for voiceCallData app
func main() {
	fmt.Println("main function for VoiceCall data app")
	voiceCallService := vcdata.GetVoiceCallService()
	bytes, err := voiceCallService.ReadCSVFile("vc.csv")
	if err != nil {
		log.Fatalln("no data")
	}
	err = voiceCallService.SetData(bytes)
	if err != nil {
		return
	}
}
