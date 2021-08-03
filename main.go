package main

import (
	"fmt"
	"github.com/jakubd/apkreport/apkreportlib"
	log "github.com/sirupsen/logrus"
	"os"
)

func setUpLogger() {
	formatter := &log.TextFormatter{
		FullTimestamp: true,
	}
	log.SetFormatter(formatter)
	log.SetLevel(log.InfoLevel)
}

func main() {
	setUpLogger()
	log.Info("running apkreport...")

	defaultCfg, err := apkreportlib.GetDefaultConfigFn()
	if err != nil {
		log.Error(err)
		os.Exit(255)
	}

	if _, err := os.Stat(defaultCfg); os.IsNotExist(err) {
		log.Info("API config file doesn't exist creating .apkreport.yml in: " + defaultCfg )
		errCfgCreate := apkreportlib.CreateDefaultConfig()
		if errCfgCreate != nil {
			log.Error("could not create config at" + defaultCfg)
			os.Exit(255)
		}
		fmt.Println("please edit config file with your API info ", defaultCfg, "and rerun")
		os.Exit(0)
	}

	outFn := "apkreport-out.csv"
	apiInfo, apiErr := apkreportlib.GetApiInfoFromConfig()

	if apiErr != nil {
		log.Error("Error parsing config file: ", apiErr)
		os.Exit(255)
	}
	log.Info("found config file at " + defaultCfg)
	log.Info("running basic report of all scans to file " + outFn)

	allRes, repErr := apkreportlib.BasicReport(apiInfo)
	if repErr != nil {
		log.Error("Error running report", repErr)
		os.Exit(255)
	}

	errCsv := apkreportlib.ApkReportSliceToCsv(allRes, outFn)
	if errCsv != nil {
		log.Error("error writing output file ", errCsv)
		os.Exit(255)
	}

	log.Info("done")
}
