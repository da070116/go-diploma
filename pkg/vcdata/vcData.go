package vcdata

import (
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"io"
	"os"
	"reflect"
	"strings"
)

type VoiceCallServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() string
}

// VoiceCallService - service to extract and store state data for VoiceCall system
type VoiceCallService struct {
	Data []VoiceCallData
}

func (v *VoiceCallService) ReadCSVFile(path string) (res []byte, err error) {
	if len(path) == 0 {
		err = errors.New("no path provided")
	}
	file, err := os.Open(path)
	if err != nil {
		return
	}

	defer utils.FileClose(file)

	res, err = io.ReadAll(file)
	if err != nil {
		return
	}
	return
}

func (v *VoiceCallService) SetData(bytes []byte) error {
	initialSize := len(v.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := v.validateData(record)
		if err != nil {
			continue
		}
		v.Data = append(v.Data, validated...)
	}
	if initialSize == len(v.Data) {
		return errors.New("no new data received")
	}
	return nil
}

func (v *VoiceCallService) ReturnData() string {
	return fmt.Sprintf("%v", v.Data)
}

func (v *VoiceCallService) validateData(record string) (validatedData []VoiceCallData, err error) {
	attrs := strings.Split(record, ";")
	if len(attrs) != reflect.TypeOf(VoiceCallData{}).NumField() {
		err = errors.New("amount of parameters provided is wrong")
		return
	}
	return
}

// GetVoiceCallService - initialize service for VoiceCall data
func GetVoiceCallService() VoiceCallServiceInterface {
	return &VoiceCallService{Data: make([]VoiceCallData, 0)}
}

// VoiceCallData - structure for store system data :
// Country - alpha-2 country code from a list
// Bandwidth - channel efficiency percent value (0 to 100)
// ResponseTime - response in milliseconds
// Provider - VoiceCall provider from a list
type VoiceCallData struct {
	Country             string
	CurrentLoading      int
	AverageResponseTime int
	Provider            string
	ConnectionClarity   float32
	CallLength          int
}
