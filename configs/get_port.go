package configs

import (
	"os"
)

func GetPort() string {
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "3000"
	}

	return Port
}
