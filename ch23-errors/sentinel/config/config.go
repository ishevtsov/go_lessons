package config

import (
	"errors"
	"os"
)

var ErrNoConfigFile = errors.New("no config file at the specified location: /tmp/myHotelApp/config.txt")

func Load() (string, error) {
	data, err := os.ReadFile("/tmp/myHotelApp/config.txt")
	if err != nil {
		return "", ErrNoConfigFile
	}
	return string(data), nil
}
