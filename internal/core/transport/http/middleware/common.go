package core_http_middleware

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	core_http_response "github.com/wander1ch/to_do_list_app/internal/core/transport/http/response"
	"go.uber.org/zap"
)

// Middleware is a function that takes an http.Handler and returns a new http.Handler


func RequestID( ) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Generate a unique request ID (you can use a library like github.com/google/uuid)
			requestID := r.Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.NewString()
			}
			
			r.Header.Set("X-Request-ID", requestID)
			w.Header().Set("X-Request-ID", requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")
			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),

			)

			ctx := core_logger.Context(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	} 
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

			
			defer func() {
				if err := recover(); err != nil {
				responseHandler.PanicResponse(err, "panic occurred while handling request")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}

func Tracing() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			
			rw := core_http_response.NewResponseWriter(w)

			befor := time.Now().UTC()

			log.Debug(
				">>> Incoming request",
				zap.String("method", r.Method),
				zap.Time("time", befor.UTC()),

			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"<<< Completed request",
				zap.Int("status_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Since(befor)),
			)
		 })
		}
}
