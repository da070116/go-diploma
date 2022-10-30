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
	DisplayData() []SupportData
	Execute(string) []SupportData
	ReturnFormattedData() []int
}

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

// SupportService - service to extract and store state data for Support system
type SupportService struct {
	Data []SupportData
}

// ReturnFormattedData - show siice [loadLevel, expectedTime] for Support data
func (s *SupportService) ReturnFormattedData() []int {
	result := make([]int, 0)
	const (
		LowSupportLoad    int = 1
		MediumSupportLoad int = 2
		HighSupportLoad   int = 3

		TicketsPerHour    int = 18
		HighLoadThreshold int = 16
		LowLoadThreshold  int = 9
	)
	activeTicketsNumber := 0
	for _, ticket := range s.Data {
		activeTicketsNumber += ticket.ActiveTickets
	}

	expectedTime := activeTicketsNumber * (60 / TicketsPerHour)
	currentSupportLoad := MediumSupportLoad

	if activeTicketsNumber >= HighLoadThreshold {
		currentSupportLoad = HighSupportLoad
	} else {
		if activeTicketsNumber < LowLoadThreshold {
			currentSupportLoad = LowSupportLoad
		}
	}

	result = append(result, currentSupportLoad, expectedTime)
	return result
}

func (s *SupportService) Execute(path string) []SupportData {
	resp, err := s.SendRequest(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = s.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return s.DisplayData()
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

// DisplayData - display Support data
func (s *SupportService) DisplayData() []SupportData {
	return s.Data
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
