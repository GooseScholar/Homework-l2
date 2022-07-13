package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"t11/internal/cache"
)

type httpServer struct {
}

func NewHttpServer(ctx context.Context, cache *cache.Cache) *httpServer {
	ServerMux := http.NewServeMux()

	ServerMux.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		
	}
	r.ParseForm()
		

		date := r.FormValue("date")
		if len(date) != 10 {
			return
		}
		newData := r.FormValue("new_data")
		if len(newData) != 10 {
			return
		}
		id := r.FormValue("user_id")
		if len(id) < 1 {
			return
		}

	})
	ServerMux.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

		date := r.FormValue("date")
		if len(date) != 10 {
			return
		}
		id := r.FormValue("user_id")
		if len(id) < 1 {
			return
		}

	})
	ServerMux.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

	})
	ServerMux.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

	})
	ServerMux.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()

	})
	go func() {
		err := http.ListenAndServe(":"+os.Getenv("portServer"), ServerMux)
		if err != nil {
			log.Println("Failed to listen" + os.Getenv("portServer"))
		}
	}()

	return &httpServer{}
}
