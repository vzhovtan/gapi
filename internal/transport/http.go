package transport

import (
	"encoding/json"
	"log"
	"net/http"

	"gapi/internal/db"
	"gapi/internal/snip"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer(snSvc *snip.Service) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /snip", func(w http.ResponseWriter, r *http.Request) {
		snippets, err := snSvc.GetAll()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		b, err := json.Marshal(snippets)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
		}
	})

	mux.HandleFunc("POST /snip", func(w http.ResponseWriter, r *http.Request) {
		var s db.Item
		err := json.NewDecoder(r.Body).Decode(&s)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = snSvc.AddItem(s)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	})

	mux.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result, err := snSvc.SearchItem(query)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		b, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		_, err = w.Write(b)
		if err != nil {
			log.Println(err)
			return
		}
	})

	return &Server{
		mux: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mux)
}
