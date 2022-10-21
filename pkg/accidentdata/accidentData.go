package accidentdata

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"go-diploma/pkg/validators"
	"io"
	"log"
	"net/http"
	"sort"
)

type AccidentServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []AccidentData
	Execute(string) []AccidentData
	ReturnFormattedData() []AccidentData
}

type AccidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

// AccidentService - service to extract and store state data for Accident system
type AccidentService struct {
	Data []AccidentData
}

type ByState []AccidentData

func (a ByState) Len() int           { return len(a) }
func (a ByState) Less(i, j int) bool { return a[i].Status < a[j].Status }
func (a ByState) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (as *AccidentService) ReturnFormattedData() []AccidentData {
	result := as.Data
	sort.Sort(ByState(result))
	return result
}

func (as *AccidentService) Execute(path string) []AccidentData {

	resp, err := as.SendRequest(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = as.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return as.DisplayData()
}

// SendRequest - function makes a GET request to provided path and returns []byte result
func (as *AccidentService) SendRequest(path string) (result []byte, err error) {
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
func (as *AccidentService) SetData(bytes []byte) error {
	initialSize := len(as.Data)

	var newRawData []AccidentData
	err := json.Unmarshal(bytes, &newRawData)
	if err != nil {
		return err
	}

	as.Data = append(as.Data, validateAccidentData(newRawData)...)
	if initialSize == len(as.Data) {
		err := errors.New("no new data")
		return err
	}
	return nil
}

// DisplayData - display Accident data
func (as *AccidentService) DisplayData() []AccidentData {
	return as.Data
}

// GetAccidentService - initialize service for Accident data
func GetAccidentService() AccidentServiceInterface {
	return &AccidentService{Data: make([]AccidentData, 0)}
}

// ValidateAccidentData - return validated Accident data from raw input
func validateAccidentData(items []AccidentData) (validItems []AccidentData) {
	for _, v := range items {
		if validators.ValidateAccidentStatus(v.Status) {
			validItems = append(validItems, v)
		}
	}
	return
}
