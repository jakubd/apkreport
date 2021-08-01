package apkreportlib

import (
	"encoding/csv"
	"os"
)

func WriteToCsvFile(givenRecords [][]string, givenFilename string) error{
	f, err := os.Create(givenFilename)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)

	err = w.WriteAll(givenRecords)
	if err != nil {
		return err
	}

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}