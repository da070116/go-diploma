package smsdata

import (
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"io"
	"os"
	"strconv"
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

// validateCountry - check whether country data is valid (belong to a list)
func validateCountry(raw string) (result string, err error) {
	if utils.IsAvailableCountry(raw) {
		result = raw
	} else {
		err = errors.New("no such country")
	}
	return
}

// validateBandwidth - check whether bandwidth data is valid (integer between 0 and 100)
func validateBandwidth(raw string) (result string, err error) {
	bandwidth, err := strconv.Atoi(raw)
	if err != nil {
		return
	}
	if bandwidth > 100 || bandwidth < 0 {
		err = errors.New("bandwidth is out of range")
		return
	}
	result = raw
	return
}

// validateResponseTime - check whether responseTime data is valid (integer)
func validateResponseTime(raw string) (result string, err error) {
	_, err = strconv.Atoi(raw)
	if err != nil {
		return
	}
	result = raw
	return
}

// validateProvider - check whether provider data is valid (string within permitted list+)
func validateProvider(raw string) (result string, err error) {
	if utils.IsAvailableProvider(raw) {
		result = raw
	} else {
		err = errors.New("no such provider")
	}
	return
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
		s.Data = append(s.Data, SMSData{
			Country:      validated[0],
			Bandwidth:    validated[1],
			ResponseTime: validated[2],
			Provider:     validated[3],
		})
	}
	if initialSize == len(s.Data) {
		return errors.New("no new data received")
	}
	fmt.Println(s.Data)
	return nil
}

// validateData - retrieve valid data array from string (if any)
func (s *SMSService) validateData(record string) (validatedData []string, err error) {
	attributes := strings.Split(record, ";")
	if len(attributes) != 4 {
		err = errors.New("amount of parameters provided is wrong")
		return
	}
	country, err := validateCountry(attributes[0])
	if err != nil {
		return
	}

	bandwidth, err := validateBandwidth(attributes[1])
	if err != nil {
		return
	}

	responseTime, err := validateResponseTime(attributes[2])
	if err != nil {
		return
	}

	provider, err := validateProvider(attributes[3])
	if err != nil {
		return
	}
	validatedData = append(validatedData, country, bandwidth, responseTime, provider)
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
