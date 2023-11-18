package util

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func GetEvnValue(key string) string {
	var foundEnvMap map[string]string = getEnvMap()

	var value string = foundEnvMap[key]

	return value
}

func getEnvMap() map[string]string {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Printf("Recover status >>> %s \n", "true")
		}
	}()

	var envMap map[string]string

	envMap, err := godotenv.Read(".env")

	if err != nil {
		log.Panicf("Found error at getting environment map >> %s", err)
	}

	defer log.Println("Defer after panic!")

	return envMap
}
