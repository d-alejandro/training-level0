package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type ModelJSONB json.RawMessage

func (modelJSONB *ModelJSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		message := fmt.Sprint("Failed to unmarshal ModelJSONB value:", value)
		return errors.New(message)
	}

	result := json.RawMessage{}
	err := json.Unmarshal(bytes, &result)
	*modelJSONB = ModelJSONB(result)
	return err
}

func (modelJSONB ModelJSONB) Value() (driver.Value, error) {
	if len(modelJSONB) == 0 {
		return nil, nil
	}
	return json.RawMessage(modelJSONB).MarshalJSON()
}
