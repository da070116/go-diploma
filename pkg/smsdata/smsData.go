package smsdata

import (
	"errors"
	"go-diploma/pkg/utils"
	"go-diploma/pkg/validators"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

type SMSServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []SMSData
	Execute(string) []SMSData
	ReturnFormattedData() [][]SMSData
}

// SMSService - service to extract and store state data for SMS system
type SMSService struct {
	Data []SMSData
}

// ReturnFormattedData - get sorted data in expected format : two sorted lists
func (s *SMSService) ReturnFormattedData() [][]SMSData {
	fullCountryData := s.displayFullCountry()
	result := make([][]SMSData, 0)
	sort.Sort(ByProviderAsc(fullCountryData))
	result = append(result, fullCountryData)

	sort.Sort(ByCountryAsc(fullCountryData))
	result = append(result, fullCountryData)
	return result
}

// Execute - initiate collector and fetch data
func (s *SMSService) Execute(filename string) []SMSData {
	path := utils.GetConfigPath(filename)
	bytes, err := s.ReadCSVFile(path)
	if err != nil {
		log.Fatalln("no data: ", err)
	}
	err = s.SetData(bytes)
	if err != nil {
		log.Fatalln("unable to set the data: ", err)
	}
	return s.DisplayData()
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

type ByCountryAsc []SMSData

func (a ByCountryAsc) Len() int           { return len(a) }
func (a ByCountryAsc) Less(i, j int) bool { return a[i].Country < a[j].Country }
func (a ByCountryAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByProviderAsc []SMSData

func (a ByProviderAsc) Len() int           { return len(a) }
func (a ByProviderAsc) Less(i, j int) bool { return a[i].Provider < a[j].Provider }
func (a ByProviderAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// SetData - append data from a file contents.
func (s *SMSService) SetData(bytes []byte) error {
	initialSize := len(s.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := s.validateData(strings.Trim(record, "\n"))
		if err != nil {
			continue
		}
		s.Data = append(s.Data, validated)
	}
	if initialSize == len(s.Data) {
		return errors.New("no new data received")
	}
	return nil
}

// validateData - retrieve valid data array from string (if any)
func (s *SMSService) validateData(record string) (validatedData SMSData, err error) {
	// important fix - string comes to function with a new-line-sign, that affects on validation
	cleanString := strings.TrimRightFunc(record, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	attrs := strings.Split(cleanString, ";")
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

func (s *SMSService) displayFullCountry() []SMSData {
	result := s.Data
	countriesMap := utils.GetCountries()
	for i, smsRecord := range s.Data {
		result[i].Country = countriesMap[smsRecord.Country]
	}
	return result
}

// DisplayData - display SMS data from service instance
func (s *SMSService) DisplayData() []SMSData {
	return s.Data
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
	res, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}
	return
}
