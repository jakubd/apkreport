package apkreportlib

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRecentScansCall(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()
	assert.NoError(t, err)

	res, err := GetRecentScansPage(apiInfo, 1)

	assert.NoError(t, err)
	assert.Greater(t, res.count, 70)
	assert.Greater(t, res.numPages, 7)
	assert.Equal(t, len(res.results), 10)
	assert.Equal(t, len(res.results[0].md5), 32)
}

func TestGetAllScans(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()
	assert.NoError(t, err)

	res, errScan := GetAllScans(apiInfo)
	assert.NoError(t, errScan)
	assert.Equal(t, len(res[0].md5), 32)
	assert.Greater(t, len(res), 69)
}

func TestGetReport(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()
	assert.NoError(t, err)

	hashImTestingWith := "01337718e5224250e9e09d645ceda74b"
	res, err := GetReport(apiInfo, hashImTestingWith)

	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(res.filename, ".apk"))
	assert.Equal(t, res.md5, hashImTestingWith)
}
