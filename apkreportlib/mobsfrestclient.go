package apkreportlib

type MobSFApiInfo struct {
	hostname string 	// include schema default: http://0.0.0.0
	port int			// default: 8000
	apiKey string		// from http://0.0.0.0:8000/api_docs
}


// http://0.0.0.0:8000/api_docs#recent-scans-api

func RecentScansCall(info *MobSFApiInfo, pageNum int) error {
	return nil
}

// http://0.0.0.0:8000/api_docs#generate-json-report-api

