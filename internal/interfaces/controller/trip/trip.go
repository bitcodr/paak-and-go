package trip

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/bitcodr/paak-and-go/internal/domain/model"
	tripservice "github.com/bitcodr/paak-and-go/internal/domain/service/trip"
	"github.com/bitcodr/paak-and-go/internal/infrastructure/helper"
)

type Controller func(http.ResponseWriter, *http.Request)

// List @Summary List of trips
// @ID list
// @Description List of trips
// @Produce json
// @Tags Trip
// @Success 200 {array} model.Trip
// @Failure 422 string http.StatusUnprocessableEntity
// @Failure 500 string http.StatusInternalServerError
// @Router /trip [get]
//.
func List(ctx context.Context, tripSrv tripservice.ITrip) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		response, err := tripSrv.List(ctx)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusUnprocessableEntity)
			return
		}

		body, err := json.Marshal(response)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			return
		}

		helper.NewResponse(&helper.Json{}).Write(w, body, http.StatusOK)
	}
}

// Show @Summary get a trip
// @ID show
// @Description get a trip
// @Produce json
// @Tags Trip
// @Success 200 {object} model.Trip
// @Failure 422 string http.StatusUnprocessableEntity
// @Failure 500 string http.StatusInternalServerError
// @Router /trip/{id} [get]
//.
func Show(ctx context.Context, tripSrv tripservice.ITrip) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var id int32

		if v, ok := vars["id"]; ok {
			parsedID, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
				return
			}

			id = int32(parsedID)
		}

		response, err := tripSrv.Show(ctx, id)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusUnprocessableEntity)
			return
		}

		body, err := json.Marshal(response)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			return
		}

		helper.NewResponse(&helper.Json{}).Write(w, body, http.StatusOK)
	}
}

// Store @Summary store a trip
// @ID store
// @Description store a trip
// @Produce json
// @Tags Trip
// @Success 200 {object} model.Trip
// @Failure 422 string http.StatusUnprocessableEntity
// @Failure 500 string http.StatusInternalServerError
// @Router /trip [post]
//.
func Store(ctx context.Context, tripSrv tripservice.ITrip) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			return
		}

		defer func() {
			if err := r.Body.Close(); err != nil {
				helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
				return
			}
		}()

		var tripModel *model.Trip

		if err := json.Unmarshal(body, &tripModel); err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			return
		}

		response, err := tripSrv.Store(ctx, tripModel)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusUnprocessableEntity)
			return
		}

		body, err = json.Marshal(response)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			return
		}

		helper.NewResponse(&helper.Json{}).Write(w, body, http.StatusOK)
	}
}
