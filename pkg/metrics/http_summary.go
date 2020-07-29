package metrics

import "net/http"

// HTTPSummary is a wrapper which observes http req/res
type HTTPSummary interface {
	Middleware(http.Handler) http.Handler
}
