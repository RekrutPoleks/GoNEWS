package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/RekrutPoleks/GoNEWS/internal/storage/psql"
	"github.com/gorilla/mux"
)

type Api struct {
	r  *mux.Router
	db *psql.DB
}

func New(db *psql.DB) *Api {
	api := Api{}
	api.r = mux.NewRouter()
	api.db = db
	api.endpoint()
	return &api
}

func (api *Api) Router() *mux.Router {
	return api.r
}

func (api *Api) endpoint() {
	api.r.HandleFunc("/news/{n}", api.News).Methods(http.MethodGet, http.MethodOptions)
	api.r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("/home/future/go/src/github.com/RekrutPoleks/GoNEWS/cmd/gonews/webapp"))))
}

func (api *Api) News(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Requset Host: %s, %s\n", r.Host, r.Method)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	s := mux.Vars(r)["n"]
	l, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	news, err := api.db.GetNews(l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err = json.NewEncoder(w).Encode(news); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
