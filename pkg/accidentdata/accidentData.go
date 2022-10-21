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
)

type AccidentServiceInterface interface {
	SendRequest(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() string
	Execute(string) string
}

type AccidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

// AccidentService - service to extract and store state data for Accident system
type AccidentService struct {
	Data []AccidentData
}

func (as *AccidentService) Execute(path string) string {

	resp, err := as.SendRequest(path)
	if err != nil {
		log.Fatalln(err)
	}
	err = as.SetData(resp)
	if err != nil {
		log.Fatalln(err)
	}
	return as.ReturnData()
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

// ReturnData - display Accident data
func (as *AccidentService) ReturnData() string {
	return fmt.Sprintf("%v\n", as.Data)
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
