package models

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"

	"github.com/pkg/errors"
)

// JSONB type based on the jinzhu/gorm one, which serialises JSON *as strings*, mostly to produce
// more useful logging than the nonsense byte arrays produce.
type JSONB []byte

func (j *JSONB) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte(`null`)) {
		*j = JSONB{}
	}
	*j = make(JSONB, len(data))
	copy(*j, data)
	return nil
}

func (j JSONB) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}

func (j JSONB) Value() (driver.Value, error) {
	if len(j) == 0 {
		return nil, nil
	}
	data, err := json.Marshal(j)
	return string(data), err
}

func (j *JSONB) Scan(v interface{}) error {
	data, ok := v.([]byte)
	if !ok {
		return errors.Errorf("scanned value for jsonb has invalid type: %T", v)
	}
	return json.Unmarshal(data, j)
}
