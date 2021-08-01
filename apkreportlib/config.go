package apkreportlib

import (
	"errors"
	"github.com/spf13/viper"
	"os"
	"os/user"
	"path/filepath"
)

// GetDefaultConfigFn - return the default config filename
func GetDefaultConfigFn() (string, error){
	userobj, err := user.Current()
	if err != nil {
		return "", err
	}

	homedir := userobj.HomeDir
	return filepath.Join(homedir, ".apkreport.yml"), nil
}

// CreateDefaultConfig - create a config file in the default dir
func CreateDefaultConfig() error{
	defaultCfg, err := GetDefaultConfigFn()
	if err != nil {
		return err
	}

	defaultConfigContents := `# MobSF API details go here.
hostname: "http://0.0.0.0"
port: 8000
apiKey: "INSERT_YOUR_KEY_HERE"
`

	if _, err := os.Stat(defaultCfg); os.IsNotExist(err) {
		f, creatErr := os.Create(defaultCfg)
		if creatErr != nil {
			return creatErr
		}

		defer f.Close()

		_, writeErr := f.WriteString(defaultConfigContents)
		if writeErr != nil {
			return writeErr
		}
	}

	if _, err := os.Stat(defaultCfg); os.IsNotExist(err) {
		return err
	}

	return nil
}

// ValidateConfig - check if all the keys are there
func ValidateConfig() error{
	mustHaveKeys := []string{"hostname", "port", "apiKey"}

	for _, thisKey := range mustHaveKeys {
		if len(viper.GetString(thisKey)) < 1 {
			return errors.New("key not in config: " + thisKey)
		}
	}
	return nil
}

// GetApiInfoFromConfig - create an api info struct based on contents of config file
func GetApiInfoFromConfig() (*MobSFApiInfo, error) {
	viper.SetConfigName(".apkreport.yml")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	// validate that all fields are there
	if err := ValidateConfig(); err != nil {
		return nil, err
	}

	// create api info struct
	apiInfoStruct := MobSFApiInfo{
		hostname: viper.GetString("hostname"),
		port: viper.GetInt("port"),
		apiKey: viper.GetString("apiKey"),
	}
	return &apiInfoStruct, nil
}
