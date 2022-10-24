package netdaemon

import (
	"fmt"
	"go-diploma/pkg/accidentdata"
	"go-diploma/pkg/billingdata"
	"go-diploma/pkg/emaildata"
	"go-diploma/pkg/mmsdata"
	"go-diploma/pkg/smsdata"
	"go-diploma/pkg/supportdata"
	"go-diploma/pkg/vcdata"
	"log"
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
	_, err := fmt.Fprintf(w, "Ok")
	if err != nil {
		log.Fatalln(err)
	}
}

func PickDataConnection(w http.ResponseWriter, r *http.Request) {
	sms := smsdata.GetSMSService()
	mms := mmsdata.GetMMSService()
	vc := vcdata.GetVoiceCallService()
	b := billingdata.GetBillingService()
	m := emaildata.GetEmailService()
	s := supportdata.GetSupportService()
	a := accidentdata.GetAccidentService()

	sms.Execute("sms.csv")
	mms.Execute("http://127.0.0.1:8383/mms")
	vc.Execute("vc.csv")
	b.Execute("conf/billing.cfg")
	m.Execute("email.csv")
	s.Execute("http://127.0.0.1:8383/support")
	a.Execute("http://127.0.0.1:8383/accendent")

	resultSet := ResultSetT{
		SMS:       sms.ReturnFormattedData(),
		MMS:       mms.ReturnFormattedData(),
		VoiceCall: vc.ReturnData(),
		Email:     m.ReturnFormattedData(),
		Billing:   b.DisplayData(),
		Support:   s.ReturnFormattedData(),
		Accidents: a.ReturnFormattedData(),
	}

	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprintf(w, "%v", resultSet)
	if err != nil {
		log.Fatalln(err)
	}
}
