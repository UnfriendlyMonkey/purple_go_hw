package req

import (
	"go/hw/4-order-api/pkg/resp"
	"net/http"
)

func HandleBody[T any](w *http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		resp.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	err = Validate(body)
	if err != nil {
		resp.Json(*w, err.Error(), http.StatusBadRequest)
		return nil, err
	}
	return &body, nil
}
