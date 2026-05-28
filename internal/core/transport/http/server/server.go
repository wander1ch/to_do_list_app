package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_http_middleware "github.com/wander1ch/to_do_list_app/internal/core/transport/http/middleware"

	core_logger "github.com/wander1ch/to_do_list_app/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux *http.ServeMux
	config Config
	log *core_logger.Logger
	middlewares []core_http_middleware.Middleware
}

func NewHTTPServer(config Config, log *core_logger.Logger, middlewares ...core_http_middleware.Middleware) *HTTPServer {
	return &HTTPServer{
		mux: http.NewServeMux(),
		config: config,
		log: log,
		middlewares: middlewares,
	}
}

func (h *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)
		
		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router),
		)
	}
}








func (h *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(h.mux, h.middlewares...)
	
	server := &http.Server{
		Addr: h.config.Addr,
		Handler: mux,
	}
	
	ch := make(chan error, 1)
	
	go func() {
		defer close(ch)

		h.log.Warn("HTTP server is starting", zap.String("addr", h.config.Addr))


		err := server.ListenAndServe()

		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("HTTP server error: %w", err)
		}
	case <-ctx.Done():
		h.log.Warn("HTTP server is shutting down")
		
		shutdownCtx, cancel := context.WithTimeout(context.Background(), h.config.ShutdownTimeout)
		defer cancel()
		
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("HTTP server shutdown error: %w", err)
		}

		h.log.Warn("HTTP server has shut down gracefully")
	}

	return nil
}

