package model

import (
	"encoding/json"
)

type Trip struct {
	ID          int32 `json:"id"`
	Origin      *City
	Destination *City
	Dates       string  `json:"dates"`
	Price       float64 `json:"price"`
}

//UnmarshalJSON custom unmarshaler that will satisfy golang Unmarshaler interface
func (t *Trip) UnmarshalJSON(in []byte) error {
	type Alias Trip

	var result = &struct {
		Origin      int32 `json:"originId"`
		Destination int32 `json:"destinationId"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	err := json.Unmarshal(in, &result)
	if err != nil {
		return err
	}

	t.Origin = &City{
		ID: result.Origin,
	}

	t.Destination = &City{
		ID: result.Destination,
	}

	return nil
}


//MarshalJSON custom unmarshaler that will satisfy golang Unmarshaler interface
func (t *Trip) MarshalJSON() ([]byte, error) {
	type Alias Trip

	var result = &struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		*Alias
	}{
		Alias: (*Alias)(t),
	}

	err := json.Marshal(&result)
	if err != nil {
		return err
	}

	t.Origin = result.Origin

	t.Destination = &City{
		ID: result.Destination,
	}

	return nil
}
