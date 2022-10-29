package emaildata

import (
	"errors"
	"go-diploma/pkg/utils"
	"go-diploma/pkg/validators"
	"io"
	"log"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

type EmailServiceInterface interface {
	ReadCSVFile(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() []EmailData
	Execute(string) map[string][][]EmailData
	ReturnFormattedData() map[string][][]EmailData
}

// EmailService - service to extract and store state data for Email system
type EmailService struct {
	Data []EmailData
}

func (es *EmailService) displayFullCountry() []EmailData {
	result := es.Data
	countriesMap := utils.GetCountries()
	for i, emailRecord := range result {
		result[i].Country = countriesMap[emailRecord.Country]
	}
	return result
}

func (es *EmailService) ReturnFormattedData() map[string][][]EmailData {
	result := make(map[string][][]EmailData)

	rawData := make(map[string][]EmailData)

	for _, value := range es.displayFullCountry() {
		rawData[value.Country] = append(rawData[value.Country], value)
	}

	for key, valuesList := range rawData {
		minTimeProviders := make([]EmailData, 0)
		minTimeProviders = append(minTimeProviders, valuesList...)
		sort.Sort(ByMinDeliveryTime(minTimeProviders))

		maxTimeProviders := make([]EmailData, 0)
		maxTimeProviders = append(maxTimeProviders, valuesList...)
		sort.Sort(ByMaxDeliveryTime(maxTimeProviders))

		if len(maxTimeProviders) > 3 {
			maxTimeProviders = maxTimeProviders[:2]
		}

		if len(minTimeProviders) > 3 {
			minTimeProviders = minTimeProviders[:2]
		}

		result[key] = append(result[key], minTimeProviders, maxTimeProviders)
	}
	return result
}

func (es *EmailService) Execute(filename string) map[string][][]EmailData {
	path := utils.GetConfigPath(filename)
	bytes, err := es.ReadCSVFile(path)
	if err != nil {
		log.Fatalln("no data: ", err)
	}
	err = es.SetData(bytes)
	if err != nil {
		log.Fatalln("unable to set data")
	}
	return es.ReturnFormattedData()
}

func (es *EmailService) ReadCSVFile(path string) (res []byte, err error) {
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

func (es *EmailService) SetData(bytes []byte) error {
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
	return nil
}

func (es *EmailService) DisplayData() []EmailData {
	return es.Data
}

func (es *EmailService) validateData(record string) (validatedData EmailData, err error) {
	// important fix - string comes to function with a new-line-sign, that affects on validation
	cleanString := strings.TrimRightFunc(record, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	attrs := strings.Split(cleanString, ";")
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

type ByMinDeliveryTime []EmailData

func (a ByMinDeliveryTime) Len() int           { return len(a) }
func (a ByMinDeliveryTime) Less(i, j int) bool { return a[i].DeliveryTime < a[j].DeliveryTime }
func (a ByMinDeliveryTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByMaxDeliveryTime []EmailData

func (a ByMaxDeliveryTime) Len() int           { return len(a) }
func (a ByMaxDeliveryTime) Less(i, j int) bool { return a[i].DeliveryTime > a[j].DeliveryTime }
func (a ByMaxDeliveryTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
