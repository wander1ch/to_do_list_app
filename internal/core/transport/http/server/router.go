package core_http_server

import (
	"net/http"
)

type ApiVersion string

var (
	APIVersionV1 ApiVersion = "v1"
	APIVersionV2 ApiVersion = "v2"
	APIVersionV3 ApiVersion = "v3"
)
// добавил routes в APIVersionRouter для хранения зарегистрированных маршрутов и их обработчиков, чтобы можно было обрабатывать несколько методов для одного пути
type APIVersionRouter struct {
	*http.ServeMux
	apiVersion ApiVersion
	routes map[string]map[string]http.HandlerFunc
}

func (r *APIVersionRouter) RegisterRoutes(UsersRoutes []Route) {
	r.RegisterRoute(UsersRoutes...)
}

func NewAPIVersionRouter(apiVersion ApiVersion) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
		routes:     make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *APIVersionRouter) RegisterRoute(routes ...Route) {
	for _, route := range routes {
		// ensure map for path
		if r.routes[route.Path] == nil {
			r.routes[route.Path] = make(map[string]http.HandlerFunc)

			// register a single mux handler for this path that dispatches by method
			path := route.Path
			r.HandleFunc(path, func(w http.ResponseWriter, req *http.Request) {
				methods := r.routes[path]
				if h, ok := methods[req.Method]; ok && h != nil {
					h(w, req)
					return
				}
				// method not allowed
				w.WriteHeader(http.StatusMethodNotAllowed)
			})
		}

		// register/overwrite handler for this method
		r.routes[route.Path][route.Method] = route.Handler
	}
}
