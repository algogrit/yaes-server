package metrics

import (
	"net/http"
	"strconv"
	"time"

	httpLogger "algogrit.com/yaes-server/pkg/http_logger"
	"github.com/prometheus/client_golang/prometheus"
)

type httpSummary struct {
	sv *prometheus.SummaryVec
}

func (h *httpSummary) Middleware(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, req *http.Request) {
		begin := time.Now()

		lw := httpLogger.NewResponseWriter(w)
		next.ServeHTTP(lw, req)

		dur := time.Since(begin).Seconds()

		labels := prometheus.Labels{
			"code":   strconv.Itoa(lw.StatusCode),
			"method": req.Method,
			"path":   req.URL.Path,
		}

		h.sv.With(labels).Observe(dur)
	}

	return http.HandlerFunc(handler)
}

// NewHTTPSummary returns a new instance of a httpSummary
func NewHTTPSummary(namespace, subsystem string) HTTPSummary {
	s := httpSummary{}

	s.sv = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Help:      "Request duration in seconds.",
		Name:      "request_duration_seconds",
		Namespace: namespace,
		Subsystem: subsystem,
	}, []string{"code", "method", "path"})

	prometheus.MustRegister(s.sv)

	return &s
}
