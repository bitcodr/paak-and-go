package helper

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	var testCreate = []struct {
		name    string
		entry   Response
		want    Response
		wantErr bool
	}{
		{
			name:    "jsonResponse",
			entry:   &Json{},
			want:    &Json{},
			wantErr: false,
		},
		{
			name:    "defaultResponse",
			entry:   nil,
			want:    &Json{},
			wantErr: false,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := NewResponse(tt.entry)

			if tt.wantErr {
				assert.Error(t, errors.New("error while init response"))
				return
			}

			switch tt.want {
			case &Json{}:
				assert.Equal(t, tt.want, got.(*Json))
			default:
				assert.Error(t, errors.New("response does not match"))
			}

		})
	}
}

func TestJson_Write(t *testing.T) {
	var testCreate = []struct {
		name    string
		body    []byte
		status  int
		want    []byte
		wantErr bool
	}{
		{
			name:    "success",
			body:    []byte(`{'id': 1, 'destination': 'Barcelona', 'origin': 'Seville', 'dates': 'Mon Sun', 'price': 43.54}`),
			status:  http.StatusOK,
			want:    []byte(`{'id': 1, 'destination': 'Barcelona', 'origin': 'Seville', 'dates': 'Mon Sun', 'price': 43.54}`),
			wantErr: false,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {

			jsonResponse := &Json{}

			recorder := httptest.NewRecorder()

			jsonResponse.Write(recorder, tt.body, tt.status)

			if tt.wantErr {
				assert.Error(t, errors.New("error while init write json response"))
				return
			}

			assert.Equal(t, tt.want, recorder.Body.Bytes())
		})
	}
}
