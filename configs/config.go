package configs

import "github.com/spf13/viper"

func LoadConfig(path string) error {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	viper.AutomaticEnv()
	return nil
}
