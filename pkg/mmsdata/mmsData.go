package mmsdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"io"
	"net/http"
)

type MMSServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() string
}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

// MMSService - service to extract and store state data for MMS system
type MMSService struct {
	Data []MMSData
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

	var newData []MMSData
	err := json.Unmarshal(bytes, &newData)
	if err != nil {
		return err
	}

	m.Data = append(m.Data, validateMMSData(newData)...)
	if initialSize == len(m.Data) {
		err := errors.New("no new data")
		return err
	}
	return nil
}

// ReturnData - display MMS data
func (m *MMSService) ReturnData() string {
	return fmt.Sprintf("%v\n", m.Data)
}

// GetMMSService - initialize service for MMS data
func GetMMSService() MMSServiceInterface {
	return &MMSService{Data: make([]MMSData, 0)}
}

// validateMMSData - return validated MMS data from raw input
func validateMMSData(items []MMSData) (validItems []MMSData) {
	for _, v := range items {
		if utils.IsAvailableProvider(v.Provider) && utils.IsAvailableCountry(v.Country) {
			validItems = append(validItems, v)
		}
	}
	return
}
