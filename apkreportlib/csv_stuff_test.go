package apkreportlib

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteToCsvFile(t *testing.T){

	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Rob", "Pike", "rob"},
		{"Ken", "Thompson", "ken"},
		{"Robert", "Griesemer", "gri"},
	}

	err := WriteToCsvFile(records, "test.csv")
	assert.NoError(t, err)

	err = os.Remove("test.csv")
	assert.NoError(t, err)
}

func TestBasicReportToCsv(t *testing.T){
	apiInfo, err := GetApiInfoFromConfig()
	assert.NoError(t, err)

	allRes, repErr := BasicReport(apiInfo)
	assert.NoError(t, repErr)
	assert.Greater(t, len(allRes), 69)

	errCsv := ApkReportSliceToCsv(allRes, "aaarep.csv")
	assert.NoError(t, errCsv)
}