package provider

import (
	"github.com/oklog/ulid"
	"github.com/tidwall/gjson"
	"math/rand"
	"time"
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

func getUniqueId() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
