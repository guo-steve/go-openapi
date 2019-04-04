package openapi

import (
	"encoding/json"
	"reflect"
)

// AdditionalProperties Object either a bool or Schema
type AdditionalProperties struct {
	Value interface{}
}

// UnmarshalJSON implements json.Unmarshaler.
func (ap *AdditionalProperties) UnmarshalJSON(data []byte) error {
	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if reflect.ValueOf(v).Kind() == reflect.Map {
		ap.Value = new(Schema)
		return json.Unmarshal(data, &ap.Value)
	}

	ap.Value = v
	return nil
}

// UnmarshalYAML implements yaml.Unmarshaler.
func (ap *AdditionalProperties) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v interface{}
	if err := unmarshal(&v); err != nil {
		return err
	}

	if reflect.ValueOf(v).Kind() == reflect.Map {
		ap.Value = new(Schema)
		return unmarshal(&ap.Value)
	}

	ap.Value = v
	return nil
}
