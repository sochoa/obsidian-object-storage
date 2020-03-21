package crud

import (
	"log"
	"net/http"
)

type loggedResponseWriter struct {
	responseWriter http.ResponseWriter
	status         int
}

func (lrw *loggedResponseWriter) Header() http.Header {
	return lrw.responseWriter.Header()
}

func (lrw *loggedResponseWriter) Write(value []byte) (int, error) {
	return lrw.responseWriter.Write(value)
}

func (lrw *loggedResponseWriter) WriteHeader(status int) {
	lrw.status = status
	lrw.responseWriter.WriteHeader(status)
}

func (lrw *loggedResponseWriter) GetStatus() int {
	return lrw.status
}

func wrapResponseWriterWithStatusLogging(w http.ResponseWriter) *loggedResponseWriter {
	return &loggedResponseWriter{
		responseWriter: w,
		status:         -1,
	}
}

func requestLoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lrw := wrapResponseWriterWithStatusLogging(w)
		handler.ServeHTTP(lrw, r)
		log.Printf("%d %s %v %s", lrw.GetStatus(), r.RemoteAddr, r.Method, r.URL)
	})
}
