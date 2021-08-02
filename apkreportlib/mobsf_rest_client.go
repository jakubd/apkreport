package apkreportlib

import (
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetAPIEndpointUrl(info *MobSFApiInfo, endpoint string) string {
	return info.hostname + ":" + strconv.Itoa(info.port) + "/api/v1/" + endpoint
}

func DoGet(url string, info *MobSFApiInfo ) (string, error) {
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

// RecentScansCall - http://0.0.0.0:8000/api_docs#recent-scans-api
func RecentScansCall(info *MobSFApiInfo, pageNum int) (*MobSFRecentScansResults ,error) {
	endpointUrl := GetAPIEndpointUrl(info, "scans")
	endpointUrl = endpointUrl + "?page=" + strconv.Itoa(pageNum)

	body, err := DoGet(endpointUrl, info)
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
			timestamp:   time.Time{},
		}
		allResults = append(allResults, thisApk)
		fmt.Println(thisApk)
		return true
	})

	count := int(gjson.Get(body, "count").Int())
	pages := int(gjson.Get(body, "num_pages").Int())

	fmt.Println(count, pages, allResults)

	finalResults := MobSFRecentScansResults{
		count: count,
		numPages: pages,
		results: allResults,
	}

	return &finalResults, nil
}

// GetReport - http://0.0.0.0:8000/api_docs#generate-json-report-api
func GetReport(info *MobSFApiInfo) error {
	u, err := url.Parse(GetAPIEndpointUrl(info, "report_json"))
	fmt.Println(u)
	return err
}

