package validators

import (
	"errors"
	"go-diploma/pkg/utils"
	"strconv"
)

// ValidateCountry - check whether country data is valid (belong to a list)
func ValidateCountry(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetCountries()) {
		result = rawValue
	} else {
		err = errors.New("no such country")
	}
	return
}

// ValidateBandwidth - check whether bandwidth data is valid (integer between 0 and 100)
func ValidateBandwidth(rawValue string) (result string, err error) {
	bandwidth, err := strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	if bandwidth > 100 || bandwidth < 0 {
		err = errors.New("bandwidth is out of range")
		return
	}
	result = rawValue
	return
}

// ValidateBandwidthAsInt - check whether bandwidth data is valid (integer between 0 and 100)
func ValidateBandwidthAsInt(rawValue string) (result int, err error) {
	bandwidth, err := strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	if bandwidth > 100 || bandwidth < 0 {
		err = errors.New("bandwidth is out of range")
		return
	}
	result = bandwidth
	return
}

// ValidateResponseTime - check whether responseTime data is valid (integer)
func ValidateResponseTime(rawValue string) (result string, err error) {
	_, err = strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	result = rawValue
	return
}

// ValidateProvider - check whether provider data is valid (string within permitted list+)
func ValidateProvider(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetProviders()) {
		result = rawValue
	} else {
		err = errors.New("no such provider")
	}
	return
}

// ValidateVoiceCallProvider - check whether provider data is valid (string within permitted list+)
func ValidateVoiceCallProvider(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetVoiceCallProviders()) {
		result = rawValue
	} else {
		err = errors.New("no such provider")
	}
	return
}

// ValidateAccidentStatus - check whether value is equal to active or closed
func ValidateAccidentStatus(statusValue string) bool {
	return utils.IsInList(statusValue, map[string]struct{}{"active": {}, "closed": {}})
}

// ValidateEmailProvider - check whether provider data is valid (string within permitted list+)
func ValidateEmailProvider(rawValue string) (result string, err error) {
	if utils.IsInList(rawValue, utils.GetEmailProviders()) {
		result = rawValue
	} else {
		err = errors.New("no such provider")
	}
	return
}

func ValidateAsInteger(rawValue string) (value int, err error) {
	value, err = strconv.Atoi(rawValue)
	if err != nil {
		return
	}
	return
}

func ValidateConnectionStability(rawValue string) (value float32, err error) {
	value64, err := strconv.ParseFloat(rawValue, 32)
	if err != nil {
		return
	}
	if value64 >= 0 && value64 <= 1 {
		value = float32(value64)
	} else {
		err = errors.New("out of range")
	}
	return
}
