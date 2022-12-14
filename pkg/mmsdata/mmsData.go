package mmsdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"io"
	"log"
	"net/http"
	"sort"
)

type MMSServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []MMSData
	Execute(string) []MMSData
	ReturnFormattedData() [][]MMSData
}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type ByCountryAsc []MMSData

func (a ByCountryAsc) Len() int           { return len(a) }
func (a ByCountryAsc) Less(i, j int) bool { return a[i].Country < a[j].Country }
func (a ByCountryAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByProviderAsc []MMSData

func (a ByProviderAsc) Len() int           { return len(a) }
func (a ByProviderAsc) Less(i, j int) bool { return a[i].Provider < a[j].Provider }
func (a ByProviderAsc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// MMSService - service to extract and store state data for MMS system
type MMSService struct {
	Data []MMSData
}

// ReturnFormattedData - return MMS data sorted by provider and by country
func (m *MMSService) ReturnFormattedData() [][]MMSData {
	fullCountryData := m.displayFullCountry()
	result := make([][]MMSData, 0)
	sort.Sort(ByProviderAsc(fullCountryData))
	result = append(result, fullCountryData)

	sort.Sort(ByCountryAsc(fullCountryData))
	result = append(result, fullCountryData)
	return result
}

// displayFullCountry - show full country name instead of alpha-2 code
func (m *MMSService) displayFullCountry() []MMSData {
	result := m.Data
	countriesMap := utils.GetCountries()
	for i, mmsRecord := range m.Data {
		result[i].Country = countriesMap[mmsRecord.Country]
	}
	return result
}

// Execute - endpoint function to collect and return related data
func (m *MMSService) Execute(s string) []MMSData {
	resp, err := m.SendRequest(s)
	if err != nil {
		log.Fatalln(err)
	}
	err = m.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return m.DisplayData()
}

// SendRequest - function makes a GET request to provided path and returns []byte result
func (m *MMSService) SendRequest(path string) (result []byte, err error) {
	response, err := http.Get(path)
	if err != nil {
		return
	}
	defer utils.CloseReader(response.Body)
	if response.StatusCode != 200 {
		err = fmt.Errorf("server at %s returns %d: %v", path, response.StatusCode, response.Status)
		return
	}
	result, err = io.ReadAll(response.Body)
	return
}

// SetData - validate and fill data in storage from a raw []byte slice
func (m *MMSService) SetData(bytes []byte) error {
	initialSize := len(m.Data)

	var newRawData []MMSData
	err := json.Unmarshal(bytes, &newRawData)
	if err != nil {
		return err
	}

	m.Data = append(m.Data, validateMMSData(newRawData)...)
	if initialSize == len(m.Data) {
		err := errors.New("no new data")
		return err
	}
	return nil
}

// DisplayData - display MMS data
func (m *MMSService) DisplayData() []MMSData {
	return m.Data
}

// GetMMSService - initialize service for MMS data
func GetMMSService() MMSServiceInterface {
	return &MMSService{Data: make([]MMSData, 0)}
}

// ValidateMMSData - return validated MMS data from raw input
func validateMMSData(items []MMSData) (validItems []MMSData) {
	for _, v := range items {
		if utils.IsInList(v.Provider, utils.GetProviders()) && utils.IsInList(v.Country, utils.GetCountries()) {
			validItems = append(validItems, v)
		}
	}
	return
}
