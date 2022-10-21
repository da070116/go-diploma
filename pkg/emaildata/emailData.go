package emaildata

import (
	"errors"
	"fmt"
	"go-diploma/pkg/utils"
	"go-diploma/pkg/validators"
	"io"
	"log"
	"os"
	"reflect"
	"strings"
)

type EmailServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	ReturnData() string
	Execute(string) string
}

// EmailService - service to extract and store state data for Email system
type EmailService struct {
	Data []EmailData
}

func (es EmailService) Execute(path string) string {
	bytes, err := es.ReadCSVFile(path)
	if err != nil {
		log.Fatalln("no data")
	}
	err = es.SetData(bytes)
	if err != nil {
		log.Fatalln("unable to set data")
	}
	return es.ReturnData()
}

func (es EmailService) ReadCSVFile(path string) (res []byte, err error) {
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

func (es EmailService) SetData(bytes []byte) error {
	initialSize := len(es.Data)
	data := string(bytes[:])
	records := strings.Split(data, "\n")
	for _, record := range records {
		validated, err := es.validateData(record)
		if err != nil {
			continue
		}
		es.Data = append(es.Data, validated)
	}
	if initialSize == len(es.Data) {
		return errors.New("no new data received")
	}
	fmt.Println(es.Data)
	return nil
}

func (es EmailService) ReturnData() string {
	return fmt.Sprintf("%v", es.Data)
}

func (es EmailService) validateData(record string) (validatedData EmailData, err error) {
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
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}
