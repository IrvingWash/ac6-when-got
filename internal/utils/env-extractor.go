package utils

import (
	"fmt"
	"strings"
)

func ExtractEnvironmentVariable(name string) (string, error) {
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