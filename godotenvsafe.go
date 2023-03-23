package godotenvsafe

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func getFirstIndexChar(str string, char rune) int {
	for i := 0; i < len(str); i++ {
		if rune(str[i]) == char {
			return i
		}
	}

	return -1
}

func parseEnvTemplate(envStr string) (map[string]string, error) {
	envMap := make(map[string]string)

	envs := strings.Split(envStr, "\n")

	for _, env := range envs {
		env = strings.Trim(env, " ")

		if env == "" || env[0:1] == "#" {
			continue
		}

		equalIdx := getFirstIndexChar(env, '=')
		if equalIdx == -1 {
			return nil, fmt.Errorf("invalid format: %s", env)
		}

		key := env[0:equalIdx]

		envMap[key] = ""
	}

	return envMap, nil
}

func Load(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return err
	}

	envFileTemplate, err := os.ReadFile(filename + ".template")
	if err != nil {
		return err
	}

	envMapTemplate, err := parseEnvTemplate(string(envFileTemplate))
	if err != nil {
		return err
	}

	missingEnvs := []string{}

	for key := range envMapTemplate {
		val := os.Getenv(key)
		if val == "" {
			missingEnvs = append(missingEnvs, key)
		}
	}

	if len(missingEnvs) == 1 {
		return fmt.Errorf("there is 1 missing environment variable: %s", missingEnvs[0])
	}

	if len(missingEnvs) > 1 {
		return fmt.Errorf("there are %d missing environment variables: %s", len(missingEnvs), strings.Join(missingEnvs, ","))
	}

	return nil
}
