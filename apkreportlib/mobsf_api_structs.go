package apkreportlib

import "time"

type MobSFApiInfo struct {
	hostname string 	// include schema default: http://0.0.0.0
	port int			// default: 8000
	apiKey string		// from http://0.0.0.0:8000/api_docs
}

type MobSFAppIndex struct {
	analyzer string
	scanType string
	filename string
	appName string
	packageName string
	versionName string
	md5 string
	timestamp time.Time
}

type MobSFRecentScansResults struct {
	count int
	numPages int
	results []MobSFAppIndex
}