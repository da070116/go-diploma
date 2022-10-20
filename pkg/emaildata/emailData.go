package emaildata

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

type EmailServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() string
}

// EmailService - service to extract and store state data for Email system
type EmailService struct {
	Data []EmailData
}

func (e EmailService) ReadCSVFile(path string) (res []byte, err error) {
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

func (e EmailService) SetData(bytes []byte) error {
	initialSize := len(e.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := e.validateData(record)
		if err != nil {
			continue
		}
		e.Data = append(e.Data, validated)
	}
	if initialSize == len(e.Data) {
		return errors.New("no new data received")
	}
	fmt.Println(e.Data)
	return nil
}

func (e EmailService) ReturnData() string {
	//TODO implement me
	panic("implement me")
}

func (e EmailService) validateData(record string) (validatedData EmailData, err error) {
	attrs := strings.Split(record, ";")
	if len(attrs) != reflect.TypeOf(EmailData{}).NumField() {
		err = errors.New("amount of parameters provided is wrong")
		return
	}

	country, err := validators.ValidateCountry(attrs[0])
	if err != nil {
		return
	}

	provider, err := validators.ValidateEmailProvider(attrs[1])
	if err != nil {
		return
	}

	deliveryTime, err := validators.ValidateAsInteger(attrs[2])
	if err != nil {
		return
	}

	validatedData = EmailData{
		Country:      country,
		Provider:     provider,
		DeliveryTime: deliveryTime,
	}

	return
}

// GetEmailService - initialize service for Email data
func GetEmailService() EmailServiceInterface {
	return &EmailService{Data: make([]EmailData, 0)}
}

// EmailData - structure for store system data :
// Country - alpha-2 country code from a list
// DeliveryTime - response in milliseconds
// Provider - Email provider from a list
type EmailData struct {
	Country      string
	Provider     string
	DeliveryTime int
}
