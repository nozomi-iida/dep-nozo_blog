package middleware

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nozomi-iida/nozo_blog/libs"
	"go.uber.org/zap"
)

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func newLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func logByStatusCode(code int) {
	switch {
	case code >= 400 && code < 600:
		libs.ZipLogger().Error("Completed",
			zap.Int("status", code),
		)
	default:
		libs.ZipLogger().Info("Completed",
			zap.Int("status", code),
		)
	}

}

// TODO: 処理時間のlogを作る
func WrapHandlerWithLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, err := ioutil.ReadAll(r.Body)
		if err != nil {
			libs.ZipLogger().Error(fmt.Sprintf("ioutil.ReadAll: %v", err))
		}
		libs.ZipLogger().Info("Started",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Any("Parameters", string(buf)),
		)

		lrw := newLoggingResponseWriter(w)

		r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		next.ServeHTTP(lrw, r)
		logByStatusCode(lrw.statusCode)
	})
}
