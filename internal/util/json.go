package util

import "encoding/json"

func PrettyJson(a any) ([]byte, error) {
	return json.MarshalIndent(a, "", "    ")
}
