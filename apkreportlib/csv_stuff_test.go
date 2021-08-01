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