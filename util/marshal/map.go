package marshal

import "encoding/json"

func MapToStruct(data any, v any) error {
	jsoned, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsoned, v)
	if err != nil {
		return err
	}

	return nil
}
