package util

import (
	"log"

	"github.com/joho/godotenv"
)

func GetEvnValue(key string) string {
	var foundEnvMap map[string]string = getEnvMap()

	var value string = foundEnvMap[key]

	return value
}

func getEnvMap() map[string]string {

	var envMap map[string]string

	envMap, err := godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Found error at getting environment map >> %s", err)
	}

	return envMap
}
