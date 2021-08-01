package apkreportlib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDefaultConfigFn(t *testing.T) {
	dir, err := GetDefaultConfigFn()
	assert.NoError(t, err)
	assert.Equal(t, "/home/jt/.apkreport.yml", dir)
}

func TestCreateDefaultConfig(t *testing.T) {
	err := CreateDefaultConfig()
	assert.NoError(t, err)
}

func TestGetApiInfoFromConfig(t *testing.T) {
	apiInfo, err := GetApiInfoFromConfig()
	assert.NoError(t, err)
	assert.Greater(t, len(apiInfo.apiKey), 2)
	assert.Greater(t, len(apiInfo.hostname), 2)
	assert.Equal(t, apiInfo.port, 8000)
}