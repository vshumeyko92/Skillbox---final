package main

import (
	"Skillbox-diploma/pkg/voicecallData"
	"fmt"
)

func main() {
	voiceCallService := voicecallData.GetVoiceCallService()
	fmt.Print(voiceCallService.Execute("voicecall.csv"), "/n")
}
