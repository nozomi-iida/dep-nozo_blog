package middleware

import (
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

func (lrw *loggingResponseWriter) WriteHeader(code int)  {
	lrw.statusCode = code	
	lrw.ResponseWriter.WriteHeader(code)
}

// TODO: リクエストボディ、処理時間のlogを作る
// TODO: status codeによってresponseのログを変更する
func WrapHandlerWithLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := make([]byte, r.ContentLength)
		r.Body.Read(body)
		libs.ZipLogger().Info("Started", 
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.Any("Parameters", string(body)),
		)

		lrw := newLoggingResponseWriter(w)

		next.ServeHTTP(lrw, r)
		libs.ZipLogger().Info("Completed", 
			zap.Int("status", lrw.statusCode),
		)
	})
}
