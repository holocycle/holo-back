package controller

import (
	"net/http"

	"github.com/holocycle/holo-back/pkg/service"
)

func ConvertToStatus(err service.Error) int {
	switch err.Code() {
	case service.NOTFOUND:
		return http.StatusNotFound

	case service.FORBIDDEN:
		return http.StatusForbidden

	case service.OUTOFRANGE:
		return http.StatusBadRequest

	case service.INTERNAL:
		fallthrough
	default:
		return http.StatusInternalServerError
	}
}
