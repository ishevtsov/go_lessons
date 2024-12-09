package config

import (
	"os"
)

func Load() (string, error) {
	data, err := os.ReadFile("/tmp/myHotelApp/config.txt")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
