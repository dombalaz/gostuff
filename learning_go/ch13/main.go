package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"time"
)

type Time struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

type Handler struct{}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accept := r.Header.Get("Accept")
	t := time.Now()
	if accept != "json" {
		w.Write([]byte(t.Format(time.RFC3339Nano)))
		return
	} else {
		ts := Time{
			DayOfWeek:  t.Weekday().String(),
			DayOfMonth: t.Day(),
			Month:      t.Month().String(),
			Year:       t.Year(),
			Hour:       t.Hour(),
			Minute:     t.Minute(),
			Second:     t.Second(),
		}
		b, err := json.Marshal(ts)
		if err != nil {
			slog.Error("error patsing time struct", "error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
	}
}

func JsonMiddleware(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("incoming request", "from", r.RemoteAddr)
		f.ServeHTTP(w, r)
	})
}

func main() {
	addr := ":8080"
	server := http.Server{
		Addr:    addr,
		Handler: JsonMiddleware(Handler{}),
	}
	slog.Info("starting http server", "address", addr)
	err := server.ListenAndServe()
	if err != nil {
		slog.Error("server error", "err", err)
	}
}
