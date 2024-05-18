package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	v0 "github.com/cloudlink-omega/backend/pkg/api/v0"
	constants "github.com/cloudlink-omega/backend/pkg/constants"
	dm "github.com/cloudlink-omega/backend/pkg/data"
)

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}

func RunServer(host string, port int, mgr *dm.Manager) {
	if mgr == nil {
		log.Fatal("[Server] Got a null data manager. This should never happen, but if you see this message it happened anyways. Aborting...")
	}

	// Init router
	r := chi.NewRouter()

	// Init DB
	mgr.InitDB()

	// Display public server nickname
	log.Printf("[Server] CLΩ Backend v%s", constants.Version)
	log.Printf("[Server] This server is called %s and is publicly available at %s", mgr.ServerNickname, mgr.PublicHostname)

	// Add logging and recovery middeware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Add Real IP middleware
	r.Use(middleware.RealIP)

	// Mount custom middleware to pass data manager into requests
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			ctx = context.WithValue(ctx, constants.DataMgrCtx, mgr)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})

	// Mount v0 route
	r.Mount("/api/v0", v0.Router)

	// Mount default route (v0)
	r.Mount("/api", v0.Router)

	// Implement static file server
	FileServer(r, "/", http.Dir("./static"))

	// Display warning on startup if authless mode is enabled
	if mgr.AuthlessMode {
		log.Printf("[Server] Authless mode is enabled - This server will accept any connection request and generate randomized auth tokens.")
	}

	// Create wait group
	var wg sync.WaitGroup

	// Start REST API
	wg.Add(1)
	go func() {
		defer wg.Done()
		StartAPI(host, port, r)
	}()

	// Wait for all services to stop
	wg.Wait()
}

func StartAPI(host string, port int, r http.Handler) error {
	err := func() error {
		// Serve root router
		log.Printf("[Server] Ready for requests at %s:%d", host, port)
		return http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), r)
	}()
	if err != nil {
		log.Fatal(err)
	}
	return err
}
