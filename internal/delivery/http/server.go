package http

import (
	"net/http"

	"skeleton/pkg/grace"

	"github.com/rs/cors"
)

// SkeletonHandler ...
type SkeletonHandler interface {
	GetSkeleton(w http.ResponseWriter, r *http.Request)
}

type UserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	server   *http.Server
	Skeleton SkeletonHandler
	User     UserHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
