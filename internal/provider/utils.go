package provider

import (
	"github.com/tidwall/gjson"
)

func getErrorFromBody(body []byte, dataPath string) string {
	errorMessages := gjson.GetBytes(body, "errors.#.message").Array()
	errorMessage := ""
	for _, e := range errorMessages {
		errorMessage = errorMessage + e.String() + " "
	}
	if errorMessage != "" {
		return errorMessage
	}
	errorMessage = gjson.GetBytes(body, dataPath+".error").String()
	if errorMessage != "" {
		return errorMessage
	}
	return ""
}
