package utils

import (
	"log"
	"os"
)

// FileClose - error-free file closing
func FileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalf("error on on %v closing: %v\n", f, err)
	}
}

// GetCountries - get all available countries
func GetCountries() map[string]string {
	return map[string]string{
		"AT": "Austria",
		"BG": "Bulgaria",
		"BL": "Saint Barthélemy",
		"CA": "Canada",
		"CH": "Switzerland",
		"DK": "Denmark",
		"ES": "Spain",
		"FR": "France",
		"GB": "United Kingdom of Great Britain and Northern Ireland",
		"MC": "Monaco",
		"NZ": "New Zealand",
		"PE": "Peru",
		"RU": "Russian Federation",
		"TR": "Turkey",
		"US": "United States of America",
	}
}

// GetSMSProviders - get available SMS Providers
func GetSMSProviders() map[string]struct{} {
	return map[string]struct{}{"Topolo": {}, "Rond": {}, "Kildy": {}}
}
