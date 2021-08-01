package apkreportlib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecentScansCall(t *testing.T) {
	apiInfo := MobSFApiInfo{
		hostname: "http://0.0.0.0",
		port: 8000,
		apiKey: "",
	}

	err := RecentScansCall(&apiInfo, 1)
	assert.NoError(t, err)
}
