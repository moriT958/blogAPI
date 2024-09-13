package middlewares

import (
	"log"
	"net/http"

	"github.com/moriT958/go-api/common"
)

// http.ResponseWriterを委譲した新しいインターフェース
type resLoggingWriter struct {
	http.ResponseWriter
	code int
}

// コンストラクタ
func NewResLoggingWriter(w http.ResponseWriter) *resLoggingWriter {
	return &resLoggingWriter{ResponseWriter: w, code: http.StatusOK}
}

// 委譲元のWriteHeaderメソッドをオーバーライド
func (rlw *resLoggingWriter) WriteHeader(code int) {
	rlw.code = code
	rlw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		traceID := newTraceID()

		log.Printf("[%d]%s %s\n", traceID, req.RequestURI, req.Method) // リクエストのログ

		ctx := common.SetTraceID(req.Context(), traceID) // リクエストのコンテキストにtraceIDを付与
		req = req.WithContext(ctx)                       // リクエストに新しいコンテキストを付与
		rlw := NewResLoggingWriter(w)
		next.ServeHTTP(rlw, req)

		log.Printf("[%d]res: %d", traceID, rlw.code) // レスポンスのログ
	})
}
