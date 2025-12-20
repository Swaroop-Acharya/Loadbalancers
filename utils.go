package utils 

import (
	"log"
	"os"
)

func HandleError(err error) {	
	log.Fatalf("Error occured: %v \n", err)
	os.Exit(1)
}
