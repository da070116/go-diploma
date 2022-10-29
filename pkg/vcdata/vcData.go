package vcdata

import (
	"errors"
	"go-diploma/pkg/utils"
	"go-diploma/pkg/validators"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

type VoiceCallServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() []VoiceCallData
	Execute(string) []VoiceCallData
}

// VoiceCallService - service to extract and store state data for VoiceCall system
type VoiceCallService struct {
	Data []VoiceCallData
}

func (vc *VoiceCallService) Execute(filename string) []VoiceCallData {
	path := utils.GetConfigPath(filename)
	bytes, err := vc.ReadCSVFile(path)
	if err != nil {
		log.Fatalln("no data: ", err)
	}
	err = vc.SetData(bytes)
	if err != nil {
		log.Fatalln("no data")
	}
	return vc.ReturnData()
}

func (vc *VoiceCallService) ReadCSVFile(path string) (res []byte, err error) {
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

func (vc *VoiceCallService) SetData(bytes []byte) error {
	initialSize := len(vc.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := vc.validateData(record)
		if err != nil {
			continue
		}
		vc.Data = append(vc.Data, validated)
	}
	if initialSize == len(vc.Data) {
		return errors.New("no new data received")
	}
	return nil
}

func (vc *VoiceCallService) ReturnData() []VoiceCallData {
	return vc.Data
}

func (vc *VoiceCallService) validateData(record string) (validatedData VoiceCallData, err error) {
	attrs := strings.Split(record, ";")
	if len(attrs) != reflect.TypeOf(VoiceCallData{}).NumField() {
		err = errors.New("amount of parameters provided is wrong")
		return
	}

	country, err := validators.ValidateCountry(attrs[0])
	if err != nil {
		return
	}

	bandwidth, err := validators.ValidateBandwidthAsInt(attrs[1])
	if err != nil {
		return
	}

	responseTime, err := validators.ValidateAsInteger(attrs[2])
	if err != nil {
		return
	}

	provider, err := validators.ValidateVoiceCallProvider(attrs[3])
	if err != nil {
		return
	}

	connectionStability, err := validators.ValidateConnectionStability(attrs[4])
	if err != nil {
		return
	}

	ttfb, err := validators.ValidateAsInteger(attrs[5])
	if err != nil {
		return
	}

	voiceClarity, err := validators.ValidateAsInteger(attrs[6])
	if err != nil {
		return
	}

	medianCallTime, err := validators.ValidateAsInteger(attrs[7])
	if err != nil {
		return
	}

	validatedData = VoiceCallData{
		Country:             country,
		Bandwidth:           bandwidth,
		ResponseTime:        responseTime,
		Provider:            provider,
		ConnectionStability: connectionStability,
		TTFB:                ttfb,
		VoiceClarity:        voiceClarity,
		MedianCallTime:      medianCallTime,
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
	Country             string  `json:"country"`
	Bandwidth           int     `json:"bandwidth"`
	ResponseTime        int     `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoiceClarity        int     `json:"voice_clarity"`
	MedianCallTime      int     `json:"median_call_time"`
}
