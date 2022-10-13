package smsdata

import (
	"errors"
	"go-diploma/pkg/utils"
	"io"
	"os"
)

type SMSServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	ParseCSVLine(line string) ([]SMSData, error)
	SetData([]byte) error
	ReturnData([]SMSData) string
}

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type SMSService struct {
	Data []SMSData
}

func (s *SMSService) SetData(bytes []byte) error {
	s
}

func GetService() SMSServiceInterface {
	return &SMSService{Data: make([]SMSData, 0)}
}

func (s *SMSService) ParseCSVLine(line string) ([]SMSData, error) {
	return s.Data, nil
}

func (s *SMSService) ReturnData(data []SMSData) string {
	return ""
}

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
