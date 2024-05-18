package v0

import (
	routes "github.com/cloudlink-omega/backend/pkg/api/v0/routes"
	dm "github.com/cloudlink-omega/backend/pkg/data"
	"github.com/go-chi/chi/v5"
)

// Create chi router
var Router = chi.NewRouter()

// Passthrough for data manager
var DataManager *dm.Manager

func init() {
	// Mount routes
	Router.Route("/", routes.RootRouter)
	Router.Route("/signaling", routes.SignalingRouter)
	Router.Route("/admin", routes.AdminRouter)
}
