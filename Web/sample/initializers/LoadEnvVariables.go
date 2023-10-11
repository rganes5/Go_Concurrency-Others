package initializers

import "sample/config"

func LoadEnvVariables() {
	config.LoadConfig()
}
