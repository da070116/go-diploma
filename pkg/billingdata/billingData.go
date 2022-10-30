package billingdata

import (
	"errors"
	"go-diploma/pkg/utils"
	"io"
	"log"
	"os"
	"strconv"
)

type BillingServiceInterface interface {
	ReadFile(path string) ([]byte, error)
	SetData([]byte) error
	DisplayData() BillingData
	Execute(string) BillingData
}

// BillingService - service to extract and store state data for Billing system
type BillingService struct {
	Data BillingData
}

// Execute - endpoint function to collect and return related data
func (bs *BillingService) Execute(filename string) (result BillingData) {
	path := utils.GetConfigPath(filename)
	bytes, err := bs.ReadFile(path)
	if err != nil {
		log.Fatalln("no data for billing service: ", err)
	}
	err = bs.SetData(bytes)
	if err != nil {
		log.Fatalln(err)
	}
	result = bs.DisplayData()
	return
}

// GetBillingService - initialize service for Billing data
func GetBillingService() BillingServiceInterface {
	return &BillingService{}
}

// BillingData - structure for store system data :
// Country - alpha-2 country code from a list
// Bandwidth - channel efficiency percent value (0 to 100)
// ResponseTime - response in milliseconds
// Provider - Billing provider from a list
type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

// SetData - append data from a file contents.
func (bs *BillingService) SetData(bytes []byte) error {
	integerMaskValue, _ := strconv.Atoi(string(bytes))
	if integerMaskValue > 255 {
		integerMaskValue = integerMaskValue / 255
	}

	flagValues := make([]bool, 6)
	sliceIndex := 0
	for i := 1; i <= 255; {
		if sliceIndex+1 > len(flagValues) {
			break
		}
		flagValues[sliceIndex] = integerMaskValue&i > 0
		i = i << 1
		sliceIndex++
	}
	bs.Data = BillingData{
		CreateCustomer: flagValues[0],
		Purchase:       flagValues[1],
		Payout:         flagValues[2],
		Recurring:      flagValues[3],
		FraudControl:   flagValues[4],
		CheckoutPage:   flagValues[5],
	}
	return nil
}

// DisplayData - display Billing data from service instance
func (bs *BillingService) DisplayData() BillingData {
	return bs.Data
}

// ReadFile - returns a byte mask from file.
func (bs *BillingService) ReadFile(path string) (res []byte, err error) {
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
