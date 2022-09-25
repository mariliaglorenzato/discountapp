package utils

import "encoding/json"

const (
	DateFormatWithoutYear = "02/01"
	DateFormat            = "02/01/2006"
)

func ToJson(data any) string {
	bytes, _ := json.MarshalIndent(data, "", "\t")
	return string(bytes)
}
