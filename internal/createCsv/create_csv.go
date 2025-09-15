package createCsv

import (
	"fmt"
	"os"

	"github.com/Bones1335/find-distance/api"
	"github.com/gocarina/gocsv"
)

func CreateCsv(distances []api.Distance) error {
	newFile, err := os.Create("calculatedDistances.csv")
	if err != nil {
		fmt.Println("error creating new file")
		return err
	}
	defer newFile.Close()

	if err := gocsv.MarshalFile(&distances, newFile); err != nil {
		fmt.Println("error marchalling csv data to new file")
		return err
	}

	return nil
}
