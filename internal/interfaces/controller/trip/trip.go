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

func List(ctx context.Context, tripSrv tripservice.ITrip) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		offset := 0
		limit := 10

		if v, ok := vars["page"]; ok {
			parsedID, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			}

			//offset starts from 0, and the page query string can be page 1, 2, 3 and so on
			if parsedID > 0 {
				offset = int(parsedID) - 1
			}
		}

		if v, ok := vars["limit"]; ok {
			parsedID, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			}

			if parsedID > 0 {
				limit = int(parsedID)
			}
		}

		response, err := tripSrv.List(ctx, offset, limit)

		body, err := json.Marshal(response)

		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
		}

		helper.NewResponse(&helper.Json{}).Write(w, body, http.StatusOK)
	}
}

func Show(ctx context.Context, tripSrv tripservice.ITrip) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		var id int32

		if v, ok := vars["id"]; ok {
			parsedID, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			}

			id = int32(parsedID)
		}

		response, err := tripSrv.Show(ctx, id)

		body, err := json.Marshal(response)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
		}

		helper.NewResponse(&helper.Json{}).Write(w, body, http.StatusOK)
	}
}

func Store(ctx context.Context, tripSrv tripservice.ITrip) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
		}

		defer func() {
			if err := r.Body.Close(); err != nil {
				helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
			}
		}()

		var tripModel *model.Trip

		if err := json.Unmarshal(body, &tripModel); err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
		}

		response, err := tripSrv.Store(ctx, tripModel)

		body, err = json.Marshal(response)
		if err != nil {
			helper.NewResponse(&helper.Json{}).Write(w, []byte(err.Error()), http.StatusInternalServerError)
		}

		helper.NewResponse(&helper.Json{}).Write(w, body, http.StatusOK)
	}
}
