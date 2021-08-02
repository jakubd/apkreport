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

	_, err = RecentScansCall(apiInfo, 8)
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
