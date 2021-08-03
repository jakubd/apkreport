package apkreportlib

func BasicReport(apiInfo *MobSFApiInfo) ([]APKReport, error) {
	scanIndex, err := GetAllScans(apiInfo)
	if err != nil {
		return nil, err
	}

	var allReports []APKReport
	for _, thisScan := range scanIndex {
		thisRep, err := GetReport(apiInfo, thisScan.md5)
		if err != nil {
			return nil, err
		}
		allReports = append(allReports, *thisRep)
	}

	return allReports, nil
}