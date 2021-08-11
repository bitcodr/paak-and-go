package trip

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	tripMock "github.com/bitcodr/paak-and-go/internal/infrastructure/mock/trip"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository/impl"
)

func TestTrip_Store(t *testing.T) {
	var testCreate = []struct {
		name      string
		entry     *model.Trip
		want      *model.Trip
		wantErr   bool
		mockError error
	}{
		{
			name: "success",
			entry: &model.Trip{
				Origin: &model.City{
					ID: 1,
				},
				Destination: &model.City{
					ID: 2,
				},
				Dates: "Sun Mon",
				Price: 32.54,
			},
			want: &model.Trip{
				ID: 1,
				Origin: &model.City{
					Name: "",
				},
				Destination: &model.City{
					Name: "",
				},
				Dates: "Sun Mon",
				Price: 32.54,
			},
			wantErr:   false,
			mockError: nil,
		},
		{
			name:      "emptyRequest",
			entry:     nil,
			want:      nil,
			wantErr:   true,
			mockError: errors.New(http.StatusText(http.StatusNoContent)),
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			c := new(tripMock.MockRepo)

			c.On("Store", ctx, tt.entry).Return(tt.want, tt.mockError)

			l := &trip{
				repo: c,
			}

			got, err := l.Store(context.Background(), tt.entry)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTrip_List(t *testing.T) {
	var testCreate = []struct {
		name      string
		want      []*model.Trip
		wantErr   bool
		mockError error
	}{
		{
			name: "success",
			want: []*model.Trip{
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
			wantErr:   false,
			mockError: nil,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			c := new(tripMock.MockRepo)

			c.On("List", ctx).Return(tt.want, tt.mockError)

			l := &trip{
				repo: c,
			}

			got, err := l.List(context.Background())

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestTrip_Show(t *testing.T) {
	var testCreate = []struct {
		name      string
		entry     int32
		want      *model.Trip
		wantErr   bool
		mockError error
	}{
		{
			name:  "success",
			entry: 1,
			want: &model.Trip{
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
			wantErr:   false,
			mockError: nil,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			c := new(tripMock.MockRepo)

			c.On("Show", ctx, tt.entry).Return(tt.want, tt.mockError)

			l := &trip{
				repo: c,
			}

			got, err := l.Show(context.Background(), tt.entry)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, got, tt.want)
		})
	}
}

func TestInitService(t *testing.T) {
	var testCreate = []struct {
		name    string
		entry   impl.TripRepo
		want    *trip
		wantErr bool
	}{
		{
			name:  "success",
			entry: impl.TripRepo(nil),
			want: &trip{
				repo: impl.TripRepo(nil),
			},
			wantErr: false,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			got := InitService(ctx, tt.entry)

			if tt.wantErr {
				assert.Error(t, errors.New("error while init service"))
				return
			}

			assert.Equal(t, tt.want, got.(*trip))
		})
	}
}
