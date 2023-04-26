package web

import (
	"Skillbox-diploma/internal/repo"
	"Skillbox-diploma/internal/struct"
	"bytes"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Router(host string, config _struct.Config, countries map[string]string) {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(structToBytes(repo.RefreshStatusPage(config, countries)))
	})
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "./web-page"))
	FileServer(r, "/", filesDir)

	http.ListenAndServe(host, r)

}

func structToBytes(s any) []byte {
	reqBodyBytes := new(bytes.Buffer)
	json.NewEncoder(reqBodyBytes).Encode(s)
	return reqBodyBytes.Bytes()
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		log.Println("FileServer не разрешает URL параметры.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rc := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rc.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
