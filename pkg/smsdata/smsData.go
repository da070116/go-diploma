package smsdata

import (
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"go-diploma/pkg/validators"
	"io"
	"os"
	"reflect"
	"strings"
)

type SMSServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() string
}

// SMSService - service to extract and store state data for SMS system
type SMSService struct {
	Data []SMSData
}

// GetSMSService - initialize service for SMS data
func GetSMSService() SMSServiceInterface {
	return &SMSService{Data: make([]SMSData, 0)}
}

// SMSData - structure for store system data :
// Country - alpha-2 country code from a list
// Bandwidth - channel efficiency percent value (0 to 100)
// ResponseTime - response in milliseconds
// Provider - SMS provider from a list
type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

// SetData - append data from a file contents.
func (s *SMSService) SetData(bytes []byte) error {
	initialSize := len(s.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := s.validateData(record)
		if err != nil {
			continue
		}
		s.Data = append(s.Data, validated)
	}
	if initialSize == len(s.Data) {
		return errors.New("no new data received")
	}
	fmt.Println(s.Data)
	return nil
}

// validateData - retrieve valid data array from string (if any)
func (s *SMSService) validateData(record string) (validatedData SMSData, err error) {
	attrs := strings.Split(record, ";")
	if len(attrs) != reflect.TypeOf(SMSData{}).NumField() {
		err = errors.New("amount of parameters provided is wrong")
		return
	}
	country, err := validators.ValidateCountry(attrs[0])
	if err != nil {
		return
	}

	bandwidth, err := validators.ValidateBandwidth(attrs[1])
	if err != nil {
		return
	}

	responseTime, err := validators.ValidateResponseTime(attrs[2])
	if err != nil {
		return
	}

	provider, err := validators.ValidateProvider(attrs[3])
	if err != nil {
		return
	}
	validatedData = SMSData{
		Country:      country,
		Bandwidth:    bandwidth,
		ResponseTime: responseTime,
		Provider:     provider,
	}
	return
}

// ReturnData - display SMS data from service instance
func (s *SMSService) ReturnData() string {
	return fmt.Sprintf("%v", s.Data)
}

// ReadCSVFile - returns a byte array from csv file.
func (s *SMSService) ReadCSVFile(path string) (res []byte, err error) {
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
