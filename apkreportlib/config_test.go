package apkreportlib

import "github.com/spf13/viper"

func CreateDefaultConfig() {

}

func GetApiInfoFromConfig() string {
	viper.SetConfigName("apkreport")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return ""
	}

	//if len(viper.GetString("maxmind_dir")) > 1 {
	//	return viper.GetString("maxmind_dir")
	//}

	return ""
}
