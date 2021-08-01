package apkreportlib

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func GetAPIEndpointUrl(info *MobSFApiInfo, endpoint string) string {
	return info.hostname + ":" + strconv.Itoa(info.port) + "/api/v1/" + endpoint
}

// RecentScansCall - http://0.0.0.0:8000/api_docs#recent-scans-api
func RecentScansCall(info *MobSFApiInfo, pageNum int) error {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", GetAPIEndpointUrl(info, "scans"), nil)
	req.Header.Set("Authorization", info.apiKey)
	res, _ := client.Do(req)

	fmt.Println(res.Body)

	return nil
}

// GetReport - http://0.0.0.0:8000/api_docs#generate-json-report-api
func GetReport(info *MobSFApiInfo) error {
	u, err := url.Parse(GetAPIEndpointUrl(info, "report_json"))
	fmt.Println(u)
	return err
}

