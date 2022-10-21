package netdaemon

import (
	"fmt"
	"go-diploma/pkg/accidentdata"
	"go-diploma/pkg/billingdata"
	"go-diploma/pkg/emaildata"
	"go-diploma/pkg/mmsdata"
	"go-diploma/pkg/smsdata"
	"go-diploma/pkg/vcdata"
	"net/http"
)

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}

type ResultSetT struct {
	SMS       [][]smsdata.SMSData                `json:"sms"`
	MMS       [][]mmsdata.MMSData                `json:"mms"`
	VoiceCall []vcdata.VoiceCallData             `json:"voice_call"`
	Email     map[string][][]emaildata.EmailData `json:"email"`
	Billing   billingdata.BillingData            `json:"billing"`
	Support   []int                              `json:"support"`
	Accidents []accidentdata.AccidentData        `json:"accident"`
}

func HandleConnection(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Ok")
}
