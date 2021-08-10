package trip

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bitcodr/paak-and-go/internal/infrastructure/repository/impl"
)

func TestInitService(t *testing.T) {
	var testCreate = []struct {
		name    string
		entry   impl.Trip
		want    *trip
		wantErr bool
	}{
		{
			name:  "nil",
			entry: nil,
			want: &trip{
				repo: nil,
			},
			wantErr: false,
		},
	}

	for _, tt := range testCreate {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.TODO()

			got := InitService(ctx, tt.entry)

			if tt.wantErr {
				assert.Error(t, errors.New("error while init service"))
				return
			}

			assert.Equal(t, tt.want, got.(*trip))
		})
	}
}

func TestTrip_List(t *testing.T) {

}

func TestTrip_Show(t *testing.T) {

}

func TestTrip_Store(t *testing.T) {

}
