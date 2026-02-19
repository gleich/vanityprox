package util

import (
	"net/http"

	"go.mattglei.ch/timber"
)

const NOT_FOUND_ERROR = "404 error: package not found"

func InternalServerError(w http.ResponseWriter, err error) {
	timber.Error(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
