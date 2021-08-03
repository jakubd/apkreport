package apkreportlib

import (
	"encoding/csv"
	"os"
)

func WriteToCsvFile(givenRecords [][]string, givenFilename string) error{
	f, err := os.Create(givenFilename)
	defer f.Close()

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

func ApkReportSliceToCsv(reports []APKReport, givenFilename string) error {
	f, err := os.Create(givenFilename)
	defer f.Close()

	if err != nil {
		return err
	}
	w := csv.NewWriter(f)
	headers := reports[0].GetBasicHeaders()
	if err := w.Write(headers); err != nil {
		return err
	}

	for _, thisReport := range reports {
		if err := w.Write(thisReport.GetBasicSlice()); err != nil {
			return err
		}
	}

	return nil
}