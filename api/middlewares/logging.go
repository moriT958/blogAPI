package middlewares

import (
	"log"
	"net/http"
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
		log.Println(req.RequestURI, req.Method) // リクエストのログ

		rlw := NewResLoggingWriter(w)
		next.ServeHTTP(rlw, req)

		log.Println("res: ", rlw.code) // レスポンスのログ
	})
}
