package riotapi

import (
	"os"
	"strings"
)

type Credentials struct {
	Port     string
	Password string
}

func ReadLockfile() (*Credentials, error) {
	lockfilePath := "C:\\Riot Games\\League of Legends\\lockfile"

	data, err := os.ReadFile(lockfilePath)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(string(data), ":")

	return &Credentials{
		Port:     parts[2],
		Password: parts[3],
	}, nil
}
