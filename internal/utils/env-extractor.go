package utils

import (
	"fmt"
	"os"
	"strings"
)

func ExtractEnvironmentVariable(name string) (string, error) {
	osVariable := os.Getenv(name)

	if osVariable != "" {
		return osVariable, nil
	}

	variables := ReadFileToArray(".env")

	var neededVariable string

	for _, v := range variables {
		if strings.HasPrefix(v, name) {
			neededVariable = strings.Split(v, "=")[1]
		}
	}

	if neededVariable == "" {
		return "", fmt.Errorf("Environment variable %s is not provided", name)
	}

	return neededVariable, nil
}
