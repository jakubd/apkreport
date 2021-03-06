package apkreportlib

import (
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetAPIEndpointUrl - construct a simple MobSF API call with the API info in the struct.
func GetAPIEndpointUrl(info *MobSFApiInfo, endpoint string) string {
	return info.hostname + ":" + strconv.Itoa(info.port) + "/api/v1/" + endpoint
}

// DoGet - Do a simple GET request and return the body as a string.
func DoGet(url string, info *MobSFApiInfo) (string, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", info.apiKey)
	res, _ := client.Do(req)

	if res.StatusCode != 200 {
		return "", errors.New("Endpoint did not return HTTP200: " + url )
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), err
}

// GetRecentScansPage - http://0.0.0.0:8000/api_docs#recent-scans-api
func GetRecentScansPage(apiInfo *MobSFApiInfo, pageNum int) (*MobSFRecentScansResults, error) {
	endpointUrl := GetAPIEndpointUrl(apiInfo, "scans")
	endpointUrl = endpointUrl + "?page=" + strconv.Itoa(pageNum)

	body, err := DoGet(endpointUrl, apiInfo)
	if err != nil {
		return nil, err
	}

	content := gjson.Get(body, "content")
	var allResults []MobSFAppIndex
	content.ForEach(func(key, value gjson.Result) bool {
		thisApkJson := value.String()
		thisApk := MobSFAppIndex{
			analyzer:    gjson.Get(thisApkJson, "ANALYZER").String(),
			scanType:    gjson.Get(thisApkJson, "SCAN_TYPE").String(),
			filename:    gjson.Get(thisApkJson, "FILE_NAME").String(),
			appName:     gjson.Get(thisApkJson, "APP_NAME").String(),
			packageName: gjson.Get(thisApkJson, "PACKAGE_NAME").String(),
			versionName: gjson.Get(thisApkJson, "VERSION_NAME").String(),
			md5:         gjson.Get(thisApkJson, "MD5").String(),
		}
		allResults = append(allResults, thisApk)
		return true
	})

	count := int(gjson.Get(body, "count").Int())
	pages := int(gjson.Get(body, "num_pages").Int())

	finalResults := MobSFRecentScansResults{
		count: count,
		numPages: pages,
		results: allResults,
	}

	return &finalResults, nil
}

// GetAllScans - return all scans info at once.
func GetAllScans(apiInfo *MobSFApiInfo) ([]MobSFAppIndex, error) {
	firstRes, err := GetRecentScansPage(apiInfo, 1)
	if err != nil {
		return nil, err
	}

	var allRes []MobSFAppIndex
	allRes = append(allRes, firstRes.results...)

	if firstRes.numPages > 1 {
		for i := 2; i < firstRes.numPages+1; i++ {
			thisPageRes, err := GetRecentScansPage(apiInfo, i)
			if err != nil {
				return nil, err
			}
			allRes = append(allRes, thisPageRes.results...)
		}
	}

	return allRes, nil
}

// GetReport - http://0.0.0.0:8000/api_docs#generate-json-report-api
func GetReport(apiInfo *MobSFApiInfo, md5Hash string) (*APKReport, error) {
	endpointUrl := GetAPIEndpointUrl(apiInfo, "report_json")
	form := url.Values{
		"hash": {md5Hash},
	}

	client := &http.Client{}
	req, _ := http.NewRequest("POST", endpointUrl, strings.NewReader(form.Encode()))
	req.Header.Set("Authorization", apiInfo.apiKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("endpoint did not return status 200 : " + endpointUrl)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	body := string(bodyBytes)
	
	report := APKReport{
		packageName:    gjson.Get(body, "package_name").String(),
		filename:       gjson.Get(body, "file_name").String(),
		md5:            gjson.Get(body, "md5").String(),
		sha1:           gjson.Get(body, "sha1").String(),
		sha256:         gjson.Get(body, "sha256").String(),
		size:           gjson.Get(body, "size").String(),
		playTitle:      gjson.Get(body, "playstore_details.title").String(),
		playDesc:       gjson.Get(body, "playstore_details.description").String(),
		playInstalls:   gjson.Get(body, "playstore_details.installs").String(),
		playDeveloper:  gjson.Get(body, "playstore_details.developer").String(),
		playDevWebsite: gjson.Get(body, "playstore_details.developerWebsite").String(),
		playDevAddress: gjson.Get(body, "playstore_details.developerAddress").String(),
		playUrl:        gjson.Get(body, "playstore_details.url").String(),
		playVersion:    gjson.Get(body, "playstore_details.version").String(),
		playPrivacyUrl: gjson.Get(body, "playstore_details.privacyPolicy").String(),
	}
	
	return &report, nil
}

