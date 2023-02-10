package util

import (
	"encoding/json"
	"fmt"
)

func PrettyJSON(a any) ([]byte, error) {
	data, err := json.MarshalIndent(a, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return data, nil
}
