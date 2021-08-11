package model

import (
	"encoding/json"
	"errors"
)

//Trip model - it's business model that has relation with city model
type Trip struct {
	ID          int32   `json:"id"`
	Origin      *City   `json:"-"` //we don't wanna the object pass to the response while encoding
	Destination *City   `json:"-"` //we don't wanna the object pass to the response while encoding
	Dates       string  `json:"dates"`
	Price       float64 `json:"price"`
}

// UnmarshalJSON custom unmarshaler that will satisfy golang Unmarshaler interface
// the reason of doing that is, we don't wanna have multiple model for request and response, we will use our
// business model and with golang custom unmarshaller can decode it to the model our domain understand it and keep
// complete association between our entities
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

//MarshalJSON custom unmarshaler that will satisfy golang Marshaler interface
// the reason of doing that is, we don't wanna have multiple model for request and response, we will use our
// business model and with golang custom marshaller can encode it to the model that our domain understand it and keep
// complete association between our entities
func (t *Trip) MarshalJSON() ([]byte, error) {
	type Alias Trip

	if t.Origin == nil || t.Destination == nil {
		return nil, errors.New("origin or destination are empty")
	}

	var result = &struct {
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		*Alias
	}{
		Origin:      t.Origin.Name,
		Destination: t.Destination.Name,
		Alias:       (*Alias)(t),
	}

	body, err := json.Marshal(&result)
	if err != nil {
		return nil, err
	}

	return body, nil
}
