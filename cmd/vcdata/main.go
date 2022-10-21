package main

import (
	"fmt"

	"go-diploma/pkg/vcdata"
)

// main function for voiceCallData app
func main() {
	fmt.Println("main function for VoiceCall data app")
	voiceCallService := vcdata.GetVoiceCallService()
	fmt.Println(voiceCallService.Execute("vc.csv"))
}
