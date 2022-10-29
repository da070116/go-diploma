package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
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

// GetProviders - get available MMS and SMS providers
func GetProviders() map[string]struct{} {
	return map[string]struct{}{"Topolo": {}, "Rond": {}, "Kildy": {}}
}

// GetVoiceCallProviders - get available VoiceCall providers
func GetVoiceCallProviders() map[string]struct{} {
	return map[string]struct{}{"TransparentCalls": {}, "E-Voice": {}, "JustPhone": {}}
}

// GetEmailProviders - return Email providers list
func GetEmailProviders() map[string]struct{} {
	return map[string]struct{}{
		"Gmail": {}, "Yahoo": {}, "Hotmail": {}, "MSN": {},
		"Orange": {}, "Comcast": {}, "AOL": {}, "Live": {}, "RediffMail": {},
		"GMX": {}, "Protonmail": {}, "Yandex": {}, "Mail.ru": {},
	}
}

// IsInList - return whether value is in list.
func IsInList[Base string | struct{}](val string, list map[string]Base) bool {
	_, ok := list[val]
	return ok
}

// GetConfigPath
func GetConfigPath(filename string) (resultPath string) {
	currentLocation, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error on calling current path: %s\n", err)
	}
	// /root/path/to/go_diploma/cmd/anyappfolder/main.go
	// os.Getwd returns             /root/path/to/go_diploma/cmd/anyappfolder
	// and config folder located at /root/path/to/go_diploma/conf
	// so we need to move two folders higher
	rootFolder := filepath.Dir(filepath.Dir(currentLocation))
	resultPath = filepath.Join(rootFolder, "conf", filename)

	fmt.Println(resultPath)

	return
}
