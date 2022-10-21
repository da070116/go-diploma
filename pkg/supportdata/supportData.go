package supportdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"io"
	"log"
	"net/http"
)

type SupportServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() string
	Execute(string) string
}

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// SupportService - service to extract and store state data for Support system
type SupportService struct {
	Data []SupportData
}

func (s *SupportService) Execute(path string) string {
	resp, err := s.SendRequest(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = s.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return s.ReturnData()
}

// SendRequest - function makes a GET request to provided path and returns []byte result
func (s *SupportService) SendRequest(path string) (result []byte, err error) {
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
func (s *SupportService) SetData(bytes []byte) error {
	initialSize := len(s.Data)

	var newRawData []SupportData
	err := json.Unmarshal(bytes, &newRawData)
	if err != nil {
		return err
	}

	s.Data = append(s.Data, validateSupportData(newRawData)...)
	if initialSize == len(s.Data) {
		err := errors.New("no new data")
		return err
	}
	return nil
}

// ReturnData - display Support data
func (s *SupportService) ReturnData() string {
	return fmt.Sprintf("%v\n", s.Data)
}

// GetSupportService - initialize service for Support data
func GetSupportService() SupportServiceInterface {
	return &SupportService{Data: make([]SupportData, 0)}
}

// ValidateSupportData - return validated Support data from raw input
func validateSupportData(items []SupportData) (validItems []SupportData) {
	for _, v := range items {
		if v.ActiveTickets >= 0 {
			validItems = append(validItems, v)
		}
	}
	return
}
