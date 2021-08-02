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

type APKReport struct {
	packageName string 				// package_name
	filename string					// file_name
	md5 string						// md5
	sha1 string						// sha1
	sha256 string					// sha256
	size string						// size
	playTitle string				// playstore_details.title
	playDesc string					// playstore_details.description
	playInstall string				// playstore_details.installs
	playDeveloper string			// playstore_details.developer
	playDevWebsite string			// playstore_details.developerWebsite
	playDevAddress string			// playstore_details.developerAddress
	playUrl string					// playstore_details.url
	playVersion string				// playstore_details.version
	playPrivacyUrl string			// playstore_details.privacyPolicy

}