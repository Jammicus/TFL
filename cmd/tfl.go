package main

import (
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
	lines.PrintStatus()
}
