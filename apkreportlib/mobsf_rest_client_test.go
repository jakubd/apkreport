package apkreportlib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecentScansCall(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()

	if err != nil {
		panic("can't read config")
	}

	err = RecentScansCall(apiInfo, 1)
	assert.NoError(t, err)
}

func TestGetReport(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()

	if err != nil {
		panic("can't read config")
	}

	err = GetReport(apiInfo)
	assert.NoError(t, err)
}
