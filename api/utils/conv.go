package utils

import (
	"net/http"
	"strconv"
)

// ConvertQueryStringToInt64 is a function
func ConvertQueryStringToInt64(r *http.Request) (int64, error) {

	query := r.URL.Query() // e.g. map[id:[1]]
	id := query.Get("id")
	return strconv.ParseInt(id, 10, 64)

}
