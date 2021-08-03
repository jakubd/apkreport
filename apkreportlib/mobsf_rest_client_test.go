package apkreportlib

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRecentScansCall(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()
	if err != nil {
		panic("can't read config")
	}

	res, err := RecentScansCall(apiInfo, 1)

	assert.NoError(t, err)
	assert.Greater(t, res.count, 70)
	assert.Greater(t, res.numPages, 7)
	assert.Equal(t, len(res.results), 10)
	assert.Equal(t, len(res.results[0].md5), 32)
}

func TestGetReport(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()

	if err != nil {
		panic("can't read config")
	}

	hashImTestingWith := "01337718e5224250e9e09d645ceda74b"
	res, err := GetReport(apiInfo, hashImTestingWith)

	assert.NoError(t, err)
	assert.True(t, strings.HasSuffix(res.filename, ".apk"))
	assert.Equal(t, res.md5, hashImTestingWith)
}
