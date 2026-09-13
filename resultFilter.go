package types

import (
	"net/http"
)

type ResultFilter func(w http.ResponseWriter, r *http.Request, result HandlerResult)
