package utils

import (
	"io"
	"log"
	"os"
)

// FileClose - error-free file closing
func FileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalf("error on %v closing: %v\n", f, err)
	}
}

// CloseReader - close reader after get request body
func CloseReader(Body io.ReadCloser) {
	err := Body.Close()
	if err != nil {
		log.Fatalln(err)
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

// GetProviders - get available Providers
func GetProviders() map[string]struct{} {
	return map[string]struct{}{"Topolo": {}, "Rond": {}, "Kildy": {}}
}

// IsAvailableCountry - return whether value is in list.
func IsAvailableCountry(val string) bool {
	list := GetCountries()
	_, ok := list[val]
	return ok
}

// IsAvailableProvider - return whether value is in list.
func IsAvailableProvider(val string) bool {
	list := GetProviders()
	_, ok := list[val]
	return ok
}
