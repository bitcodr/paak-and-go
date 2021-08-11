package trip

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	tripMock "github.com/bitcodr/paak-and-go/internal/infrastructure/mock/trip"
)

func TestList(t *testing.T) {
	var testCreate = []struct {
		name      string
		entry     []*model.Trip
		want      []byte
		wantErr   bool
		mockError error
	}{
		{
			name: "success",
			entry: []*model.Trip{
				{
					ID: 1,
					Origin: &model.City{
						Name: "Barcelona",
					},
					Destination: &model.City{
						Name: "Valencia",
					},
					Dates: "Sun Mon",
					Price: 32.54,
				},
				{
					ID: 2,
					Origin: &model.City{
						Name: "Andorra la Vella",
					},
					Destination: &model.City{
						Name: "Malaga",
					},
					Dates: "Thu Wed",
					Price: 65.9,
				},
			},
			want:      []byte(`[{"origin":"Barcelona","destination":"Valencia","id":1,"dates":"Sun Mon","price":32.54},{"origin":"Andorra la Vella","destination":"Malaga","id":2,"dates":"Thu Wed","price":65.9}]`),
			wantErr:   false,
			mockError: nil,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			c := new(tripMock.MockService)

			c.On("List", ctx).Return(tt.entry, tt.mockError)

			got := List(ctx, c)

			if tt.wantErr {
				assert.Error(t, errors.New("error while init controller"))
				return
			}

			recorder := httptest.NewRecorder()

			request := httptest.NewRequest(http.MethodGet, "/trip", nil)

			got(recorder, request)

			assert.Equal(t, tt.want, recorder.Body.Bytes())
		})
	}
}

func TestShow(t *testing.T) {
	var testCreate = []struct {
		name      string
		id        int32
		entry     *model.Trip
		want      []byte
		wantErr   bool
		mockError error
	}{
		{
			name: "success",
			id:   1,
			entry: &model.Trip{
				ID: 1,
				Origin: &model.City{
					Name: "Barcelona",
				},
				Destination: &model.City{
					Name: "Valencia",
				},
				Dates: "Sun Mon",
				Price: 32.54,
			},
			want:      []byte(`{"origin":"Barcelona","destination":"Valencia","id":1,"dates":"Sun Mon","price":32.54}`),
			wantErr:   false,
			mockError: nil,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			c := new(tripMock.MockService)

			c.On("Show", ctx, tt.id).Return(tt.entry, tt.mockError)

			got := Show(ctx, c)

			if tt.wantErr {
				assert.Error(t, errors.New("error while init controller"))
				return
			}

			recorder := httptest.NewRecorder()

			request := httptest.NewRequest(http.MethodGet, "/trip/"+strconv.Itoa(int(tt.id)), nil)

			urlParams := map[string]string{
				"id": strconv.Itoa(int(tt.id)),
			}

			request = mux.SetURLVars(request, urlParams)

			got(recorder, request)

			assert.Equal(t, tt.want, recorder.Body.Bytes())
		})
	}
}

func TestStore(t *testing.T) {

}
