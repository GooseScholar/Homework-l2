package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"t11/internal/cache"
	"t11/internal/handler"
	"t11/internal/middleware"
)

type httpServer struct {
}

//NewHTTPServer запуск http сервера
func NewHTTPServer(ctx context.Context, cache *cache.Cache) {
	ServerMux := http.NewServeMux()

	ServerMux.HandleFunc("/create_event", func(w http.ResponseWriter, r *http.Request) {
		handler.CreateEvent(w, r, cache)
	})
	ServerMux.HandleFunc("/update_event", func(w http.ResponseWriter, r *http.Request) {
		handler.UpdateEvent(w, r, cache)
	})
	ServerMux.HandleFunc("/delete_event", func(w http.ResponseWriter, r *http.Request) {
		handler.DeleteEvent(w, r, cache)
	})
	ServerMux.HandleFunc("/events_for_day", func(w http.ResponseWriter, r *http.Request) {
		handler.EventsForDay(w, r, cache)
	})
	ServerMux.HandleFunc("/events_for_week", func(w http.ResponseWriter, r *http.Request) {
		handler.EventsForWeek(w, r, cache)
	})
	ServerMux.HandleFunc("/events_for_month", func(w http.ResponseWriter, r *http.Request) {
		handler.EventsForMonth(w, r, cache)
	})

	go func() {
		err := http.ListenAndServe(":"+os.Getenv("portServer"), middleware.Logging(ServerMux))
		if err != nil {
			log.Println("Failed to listen" + os.Getenv("portServer"))
		}
	}()

	return
}
