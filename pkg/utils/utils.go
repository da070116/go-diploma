package utils

import (
	"fmt"
	"log"
	"os"
)

func FileClose(f *os.File) {
	err := f.Close()
	if err != nil {
		log.Fatalln(fmt.Sprintf("error on on %v closing: %v\n", f, err))
	}
}
