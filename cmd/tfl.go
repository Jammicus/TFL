package main

import (
	"fmt"
	"os"
	lines "tfl/internal"
)

func getEnvironmentVariables() map[string]string {
	envVars := make(map[string]string)
	envVars["applicationID"] = os.Getenv("APPID")
	envVars["applicationKeys"] = os.Getenv("APPKEY")
	return envVars
}

func main() {

	envVars := getEnvironmentVariables()
	fmt.Println(envVars["applicationID"])
	fmt.Println(envVars["applicationKeys"])
	lines.PrintStatus()
}
