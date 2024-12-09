package convert

import "encoding/json"

func ToStruct[T any](d any) (T, error) {
	var res T

	mar, err := json.Marshal(&d)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(mar, &res); err != nil {
		return res, err
	}

	return res, nil
}
